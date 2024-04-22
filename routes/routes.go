package routes

import (
	"greet-home-srv/configs"

	"github.com/greetinc/greet-middlewares/middlewares"
	"github.com/labstack/echo/v4"

	h_user "greet-home-srv/handlers"
	r_user "greet-home-srv/repositories"
	s_user "greet-home-srv/services"
)

var (
	DB = configs.InitDB()

	JWT   = middlewares.NewJWTService()
	userR = r_user.NewUserRepository(DB)
	userS = s_user.NewUserService(userR, JWT)
	userH = h_user.NewUserHandler(userS)
)

func New() *echo.Echo {

	e := echo.New()
	v1 := e.Group("api/v1")
	{
		user := v1.Group("/user", middlewares.AuthorizeJWT(JWT))
		{
			user.GET("/find", userH.GetAll)
		}
	}
	return e
}
