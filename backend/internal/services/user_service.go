package services

import (
	"errors"
	"time"

	"github.com/DB-Vincent/want-to-read/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

var sampleUser = models.User{
	ID:       1,
	Username: "vincent",
	Password: "password",
}

var jwtKey = []byte("s0m3_s3cr3t_k3y")

func (s *UserService) Authenticate(username, password string) (*models.User, error) {
	if username == sampleUser.Username && password == sampleUser.Password {
		return &sampleUser, nil
	}
	return nil, errors.New("invalid credentials")
}

func (s *UserService) GenerateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func (s *UserService) ParseJWT(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})
}
