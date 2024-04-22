package repositories

import (
	"greet-like-srv/dto"
	"greet-like-srv/entity"

	entityuser "github.com/greetinc/greet-auth-srv/entity"
)

func (r *likeRepository) Get(req dto.LikeRequest) ([]dto.LikeResponse, error) {
	var data []entity.LikeRelation

	err := r.DB.Where("like_id =?", req.UserID).Model(&entity.LikeRelation{}).Find(&data).Error
	if err != nil {
		return []dto.LikeResponse{}, err
	}

	var responses []dto.LikeResponse
	for _, likeRelation := range data {
		// Periksa nilai IsFollowing, dan jika true, skip entitas ini
		if likeRelation.IsFollowing {
			continue
		}

		response := dto.LikeResponse{
			ID:     likeRelation.ID,
			UserID: likeRelation.UserID,
			LikeID: likeRelation.LikeID,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (r *likeRepository) GetByUserID(userID string) (entityuser.UserDetail, error) {
	var userDetail entityuser.UserDetail
	err := r.DB.Where("user_id = ?", userID).First(&userDetail).Error
	if err != nil {
		return entityuser.UserDetail{}, err
	}

	return userDetail, nil
}
