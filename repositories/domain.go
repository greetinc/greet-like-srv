package repositories

import (
	"greet-like-srv/dto"

	entity "github.com/greetinc/greet-auth-srv/entity"

	"gorm.io/gorm"
)

type LikeRepository interface {
	Get(req dto.LikeRequest) ([]dto.LikeResponse, error)
	GetByUserID(userID string) (entity.UserDetail, error)
}

type likeRepository struct {
	DB *gorm.DB
}

func NewLikeRepository(DB *gorm.DB) LikeRepository {
	return &likeRepository{
		DB: DB,
	}
}
