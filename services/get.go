package services

import (
	"greet-like-srv/dto"
)

func (s *likeService) Get(req dto.LikeRequest) ([]dto.LikeResponse, error) {
	data, err := s.LikeR.Get(req)
	if err != nil {
		return []dto.LikeResponse{}, err
	}

	var responses []dto.LikeResponse
	for _, likeRelation := range data {
		// Dapatkan detail pengguna berdasarkan UserID
		userDetail, err := s.LikeR.GetByUserID(likeRelation.UserID)
		if err != nil {
			return []dto.LikeResponse{}, err
		}

		response := dto.LikeResponse{
			ID:       likeRelation.ID,
			UserID:   likeRelation.UserID,
			LikeID:   likeRelation.LikeID,
			FullName: userDetail.FullName,
		}
		responses = append(responses, response)
	}

	return responses, nil
}
