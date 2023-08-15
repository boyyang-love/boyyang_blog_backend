package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Likes struct {
	Id           uint       `json:"id" gorm:"primary_key"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	Created      int        `json:"created" gorm:"autoCreateTime"`
	Updated      int        `json:"updated" gorm:"autoUpdateTime"`
	ExhibitionId uint       `json:"exhibition_id"` // 收藏图片 id
	UserId       uint       `json:"user_id"`       // 收藏者 id
	LikesType    bool       `json:"likes_type"`    // true 收藏 false 取消收藏
}

func (likes *Likes) TableName() string {
	return "likes"
}

func (likes *Likes) BeforeCreate(db *gorm.DB) (err error) {
	likes.Id = uint(uuid.New().ID())
	return
}
