package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Star struct {
	Id        uint       `json:"id" gorm:"primary_key"`
	Uid       uint32     `json:"uid" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Created   int        `json:"created" gorm:"autoCreateTime"`
	Updated   int        `json:"updated" gorm:"autoUpdateTime"`
	UserId    uint32     `json:"user_id"`
	StarId    uint32     `json:"star_id"`
	StarType  bool       `json:"star_type"` // 0 取消star 1 star
	Type      int        `json:"type"`      // 1 图片 2 博客
}

func (star *Star) TableName() string {
	return "star"
}

func (star *Star) BeforeCreate(*gorm.DB) (err error) {
	uid, err := uuid.NewUUID()
	star.Uid = uid.ID()
	return err
}
