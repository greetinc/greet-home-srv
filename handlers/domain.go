package handlers

import (
	s "greet-home-srv/services"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	GetAll(c echo.Context) error
}

type domainHandler struct {
	domainService s.UserService
}

func NewUserHandler(UserS s.UserService) UserHandler {
	return &domainHandler{
		domainService: UserS,
	}
}
