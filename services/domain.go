package services

import (
	dto "greet-home-srv/dto/user"
	m "greet-home-srv/middlewares"
	repositories "greet-home-srv/repositories/user"
)

type UserService interface {
	GetAll(dto.UserRequest) ([]dto.UserResponse, error)
	GetFriend(req dto.FriendRequest) ([]dto.FriendResponse, error)
	Create(req dto.LikeRequest) (dto.LikeResponse, error)
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
