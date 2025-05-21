package seed

import (
	"github.com/DB-Vincent/want-to-read/internal/models"
	"github.com/DB-Vincent/want-to-read/internal/services"
	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

func All() []Seed {
	return []Seed{
		{
			Name: "Create admin user",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "admin", "P@ssw0rd", true)
			},
		},
	}
}

func CreateUser(db *gorm.DB, username, password string, isSuper bool) error {
	user := &models.User{
		Username: username,
		Password: password,
		IsSuper:  isSuper,
	}

	userService := services.NewUserService()
	hashedPassword, err := userService.GenerateHash(password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	if err := db.FirstOrCreate(user).Error; err != nil {
		return err
	}
	return nil
}
