package users

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserRequestBody struct {
	Username string `json:"username" bindings:"required,alphanum,min=3,max=20"`
	Email    string `json:"email" bindings:"required,email"`
	Password string `json:"password" bindings:"required,min=8"`
}

type UserResponseBody struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func toResponse(u *User) UserResponseBody {
	return UserResponseBody{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
}

func (handler *Handler) HandleCreateUser(c *gin.Context, username, email, password string) {
	var userInput UserRequestBody
	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request."})
		return
	}

	//user, err := handler.service.CreateNewUser(c.Request.Context(), username, email, password)
	//if err != nil {
	//}
}
