package services

import (
	dto "greet-home-srv/dto/user"
	"greet-home-srv/repositories"

	m "github.com/greetinc/greet-middlewares/middlewares"
)

type UserService interface {
	GetAll(dto.UserRequest) ([]dto.UserResponse, error)
	IsValidProfileID(profileID string) bool
}

type userService struct {
	UserR repositories.UserRepository
	jwt   m.JWTService
}

func NewUserService(UserR repositories.UserRepository, jwtS m.JWTService) UserService {
	return &userService{
		UserR: UserR,
		jwt:   jwtS,
	}
}
