package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Likes struct {
	Id        uint       `json:"id" gorm:"primary_key"`
	Uid       uint       `json:"uid" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Created   int        `json:"created" gorm:"autoCreateTime"`
	Updated   int        `json:"updated" gorm:"autoUpdateTime"`
	UserId    uint       `json:"user_id"` // 收藏者 id
	LikesId   uint       `json:"likes_id"`
	LikesType bool       `json:"likes_type"` // true 收藏 false 取消收藏
	Type      int        `json:"type"`       // 1 图片 2 博客
}

func (likes *Likes) TableName() string {
	return "likes"
}

func (likes *Likes) BeforeCreate(*gorm.DB) (err error) {
	uid, err := uuid.NewUUID()
	likes.Uid = uint(uid.ID())
	return err
}
