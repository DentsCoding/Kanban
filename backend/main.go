package main

//TODO(): setup logging in, sessions, redis server

import (
	"crypto/rand"
	"crypto/subtle"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/argon2"
)

type UserIngest struct {
	Username string `json:"username" binding:"required,alphanum,min=5,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type ArgonParams struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type LoginRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required"`
}

var DefaultParams = ArgonParams{
	Memory:      64 * 1024,
	Iterations:  1,
	Parallelism: 4,
	SaltLength:  16,
	KeyLength:   32,
}

func LoginHandler(c *gin.Context) {

	var input LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	var user struct {
		ID       int
		Username string
		Password string
	}
	// 1. get username
	query := "SELECT id, username, password FROM users WHERE username = $1"
	err := db.QueryRow(query, input.Username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password."})
		return
	}
	// 2. verify password
	match, err := VerifyPassword(input.Password, user.Password)
	if err != nil || !match {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or passworod."})
		return
	}
	// 3. generate token
	tokenString, err := GenerateJWT(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token."})
		return
	}
	// 4. set cookie with said token
	c.SetCookie(
		"token",
		tokenString,
		86400, // should probably be taken from the jwt or from a source that also sets the jwt time
		"/",
		"",
		false, // only send over https, set to true in prod
		true,
	)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful."})
}

func HashFunction(password string, p ArgonParams) (string, error) {

	salt := make([]byte, p.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	b64salt := base64.RawStdEncoding.EncodeToString(salt)
	b64hash := base64.RawStdEncoding.EncodeToString(hash)

	phcFormatted := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, p.Memory, p.Iterations, p.Parallelism, b64salt, b64hash)

	return phcFormatted, nil

}

func LogoutHandler(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Log out successful."})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the jwt token
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication cookie required"})
			c.Abort()
			return
		}

		// parse and validate
		claims, err := ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalid or expired."})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()

	}
}

func GenerateJWT(userID int, username string) (string, error) {

	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	// 1. gen jwt with claims + choose sign method 2. return signed jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	// parse
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return jwtSecret, nil // function returns the secret key
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("Invalid token.")
	}
	// return
	return claims, nil
}

func VerifyPassword(password, encodedHash string) (bool, error) {

	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 || parts[1] != "argon2id" {
		return false, errors.New("Invalid hash format")
	}

	var version int
	_, err := fmt.Sscanf(parts[2], "v=%d", &version)
	if err != nil {
		return false, errors.New("Incompatible argon2 version")
	}

	var params ArgonParams
	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &params.Memory, &params.Iterations, &params.Parallelism)
	if err != nil {
		return false, errors.New("invalid parameter format")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	params.KeyLength = uint32(len(decodedHash))

	candidateHash := argon2.IDKey([]byte(password), salt, params.Iterations, params.Memory, params.Parallelism, params.KeyLength)

	// 3. Constant-time comparison to prevent timing attacks
	if subtle.ConstantTimeCompare(decodedHash, candidateHash) == 1 {
		return true, nil
	}

	return false, nil
}

func CreateUser(c *gin.Context) {
	var input UserIngest
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing credentials"})
		return
	}

	var username string
	err := db.QueryRow("SELECT username FROM users WHERE username = $1", input.Username).Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing credentials"})
			return
		}
	}

	hashedPassword, err := HashFunction(input.Password, DefaultParams)

	insertUserQuery := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id, username, email, created_at"
	var resp UserResponse
	err = db.QueryRow(insertUserQuery, input.Username, input.Email, hashedPassword).Scan(&resp.ID, &resp.Username, &resp.Email, &resp.CreatedAt)
	if err != nil {
		log.Printf("Failed to insert user into database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"}) // probably should make it so error is similar to bad credentials
		return
	}

	c.JSON(http.StatusCreated, resp)
}

var db *sql.DB

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	r := gin.Default()
	connStr := "host=localhost user=devuser password=devpassword123 database=dev_database sslmode=disable"
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("Error connecting to the database")
	}

	defer db.Close()

	r.POST("/api/users/create", CreateUser)
	r.POST("/api/auth/login", LoginHandler)
	r.POST("/api/auth/logout", LogoutHandler)

	protected := r.Group("/api/auth")
	protected.Use(AuthMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			userID, _ := c.Get("userID")
			username, _ := c.Get("username")

			c.JSON(http.StatusOK, gin.H{
				"user_id":  userID,
				"username": username,
				"status":   "authenticated",
			})
		})

	}

	r.Run(":5000")

}
