package services

import (
	"context"
	"errors"
	"os"
	"time"

	"example.com/m/v2/models"
	"example.com/m/v2/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

var ErrInvalidCredentials = errors.New("invalid username or password")

func LoginAdmin(ctx context.Context, payload models.LoginPayload) (string, error) {
	admin, err := repository.GetAdminByUsername(ctx, payload.Username)
	if err != nil {
		return "", err
	}
	if admin == nil {
		return "", ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password_hash), []byte(payload.Password))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": admin.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
