package auth

import (
	"fmt"
	"time"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type JWTKey = []byte

const passwordMinLength = 8

func HashPassword(password string) (string, error) {
	if len(password) < passwordMinLength {
		return "", fmt.Errorf("password must be at least %d characters long", passwordMinLength)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedPassword), nil
}

func CheckPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func GenerateJWTKey(user domain.User, key JWTKey, exp time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.ID().String(),
		"email": user.Email().String(),
		"exp":   time.Now().Add(exp).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
