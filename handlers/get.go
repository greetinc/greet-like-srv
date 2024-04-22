package handlers

import (
	"greet-like-srv/dto"

	res "github.com/greetinc/greet-util/s/response"

	"github.com/labstack/echo/v4"
)

func (b *domainHandler) Get(c echo.Context) error {
	var req dto.LikeRequest

	createdBy, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	req.UserID = createdBy

	get, err := b.domainService.Get(req)
	if err != nil {
		return c.JSON(500, "Error fetching get")
	}

	return c.JSON(200, get)
}
