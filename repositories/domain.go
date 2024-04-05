package repositories

import (
	dto "github.com/greetinc/greet-auth-srv/dto/auth"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll(req dto.UserRequest) ([]dto.UserResponse, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userRepository{
		DB: DB,
	}
}
