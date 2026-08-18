package pkg

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	Memory = 64 * 1024
	Iterations = 3
	Parallelism = 2
	SaltLength = 16
	KeyLength = 32
)


func HashPassword(plainPassword string) (string, error) {
	salt := make([]byte, SaltLength)
	
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(plainPassword), salt, Iterations, Memory, Parallelism, KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	hashedPassword := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, Memory, Iterations, Parallelism, b64Salt, b64Hash,
	)

	return hashedPassword, nil
}


func VerifyPassword(plainPassword, storedPassword string) (bool, error) {
	parts := strings.Split(storedPassword, "$")
	if len(parts) != 6 {
		return false, errors.New("Invalid hash format.")
	}
	// check version

	var version int
	var memory, iterations uint32
	var parallelism uint8

	if _, err := fmt.Sscanf(parts[2], "v=%d", &version); err != nil {
		return false, errors.New("Invalid hasher version.")
	}

	if _, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism); err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, errors.New("Invalid salt hash.")
	}

	storedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, errors.New("Invalid password hash.")
	}

	newHash := argon2.IDKey([]byte(plainPassword), []byte(salt), iterations, memory, parallelism, uint32(len(storedHash)))

	if subtle.ConstantTimeCompare(storedHash, newHash) == 1 {
		return true, nil
	}

	return false, nil
}