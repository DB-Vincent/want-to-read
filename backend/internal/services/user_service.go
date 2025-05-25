package services

import (
	"errors"
	"log"
	"time"

	"github.com/DB-Vincent/want-to-read/internal/database"
	"github.com/DB-Vincent/want-to-read/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

var jwtKey = []byte("s0m3_s3cr3t_k3y")

func (s *UserService) Authenticate(user *models.User) (*models.User, error) {
	var dbUser models.User

	if user.Username == "" || user.Password == "" {
		return nil, errors.New("username and password are required")
	}

	if err := database.DB.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		return nil, errors.New("username or password is incorrect")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		return nil, errors.New("username or password is incorrect")
	}

	return &dbUser, nil
}

func (s *UserService) Register(user *models.User) (*models.User, error) {
	hashedPassword, err := s.GenerateHash(user.Password)
	user.Password = hashedPassword
	if err != nil {
		return nil, err
	}

	result := database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (s *UserService) ChangePassword(user *models.User) error {
	if user.Password == "" {
		return errors.New("password is required")
	}

	hashedPassword, err := s.GenerateHash(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	result := database.DB.Save(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) ListUsers() ([]models.User, error) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) UpdateUser(user *models.User) (*models.User, error) {
	if err := database.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s *UserService) GenerateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"is_super": user.IsSuper,
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

func (s *UserService) GetUserId(token string) (uint, error) {
	parsedToken, err := s.ParseJWT(token)
	if err != nil {
		return 0, err
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		if userId, ok := claims["user_id"].(float64); ok {
			log.Print("User id: ", userId)

			return uint(userId), nil
		}
		return 0, errors.New("user_id not found in token claims")
	}
	return 0, errors.New("invalid token claims")
}

func (s *UserService) IsSuperUser(token string) (bool, error) {
	parsedToken, err := s.ParseJWT(token)
	if err != nil {
		return false, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		if isSuper, ok := claims["is_super"].(bool); ok {
			log.Print("User is super: ", isSuper)

			return bool(isSuper), nil
		}

		return false, errors.New("is_super not found in token claims")
	}

	return false, errors.New("invalid token claims")
}
