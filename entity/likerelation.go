package entity

import (
	entity "github.com/greetinc/greet-auth-srv/entity"
)

type LikeRelation struct {
	ID          string      `gorm:"primary_key" json:"id"`
	UserID      string      `gorm:"type:varchar(36);index" json:"user_id"`
	LikeID      string      `gorm:"like_id" json:"like_id"`
	IsFollowing bool        `gorm:"is_following" json:"is_following"`
	User        entity.User `gorm:"foreignKey:UserID" json:"user"`
}
