package repositories

import (
	dto "greet-home-srv/dto/user"

	entity "github.com/greetinc/greet-auth-srv/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll(req dto.UserRequest) ([]dto.UserResponse, error)
	GetUserCoordinates(userID string) (entity.RadiusRange, error)
	GetUserByProfileID(profileID string) (*entity.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userRepository{
		DB: DB,
	}
}
