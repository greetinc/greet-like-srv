package services

import (
	"greet-like-srv/dto"
	"greet-like-srv/repositories"

	m "github.com/greetinc/greet-middlewares/middlewares"
)

type LikeService interface {
	Get(req dto.LikeRequest) ([]dto.LikeResponse, error)
}

type likeService struct {
	LikeR repositories.LikeRepository
	jwt   m.JWTService
}

func NewLikeService(LikeR repositories.LikeRepository, jwtS m.JWTService) LikeService {
	return &likeService{
		LikeR: LikeR,
		jwt:   jwtS,
	}
}
