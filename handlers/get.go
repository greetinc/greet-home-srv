package handlers

import (
	"fmt"

	dto "greet-home-srv/dto/user"

	e "github.com/greetinc/greet-util/s/response"

	"github.com/labstack/echo/v4"
)

func (b *domainHandler) GetAll(c echo.Context) error {
	var req dto.UserRequest

	userId, ok := c.Get("UserId").(string)
	if !ok {
		return e.ErrorBuilder(&e.ErrorConstant.InternalServerError, nil).Send(c)
	}

	req.ID = userId
	fmt.Println(req.ID)
	fmt.Println(userId)

	users, err := b.domainService.GetAll(req)
	if err != nil {
		return c.JSON(500, "Error fetching users")
	}

	return c.JSON(200, users)
}
