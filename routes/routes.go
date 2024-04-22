package routes

import (
	"greet-like-srv/configs"

	"github.com/greetinc/greet-middlewares/middlewares"
	"github.com/labstack/echo/v4"

	h_like "greet-like-srv/handlers"
	r_like "greet-like-srv/repositories"
	s_like "greet-like-srv/services"
)

var (
	DB = configs.InitDB()

	JWT   = middlewares.NewJWTService()
	likeR = r_like.NewLikeRepository(DB)
	likeS = s_like.NewLikeService(likeR, JWT)
	likeH = h_like.NewLikeHandler(likeS)
)

func New() *echo.Echo {

	e := echo.New()
	v1 := e.Group("api/v1")
	{
		like := v1.Group("/like", middlewares.AuthorizeJWT(JWT))
		{
			like.GET("", likeH.Get)
		}
	}
	return e
}
