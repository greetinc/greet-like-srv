package handlers

import (
	s "greet-like-srv/services"

	"github.com/labstack/echo/v4"
)

type LikeHandler interface {
	Get(c echo.Context) error
	// GetById(c echo.Context) error
	// Update(c echo.Context) error
	// Delete(c echo.Context) error
}

type domainHandler struct {
	domainService s.LikeService
}

func NewLikeHandler(LikeS s.LikeService) LikeHandler {
	return &domainHandler{
		domainService: LikeS,
	}
}
