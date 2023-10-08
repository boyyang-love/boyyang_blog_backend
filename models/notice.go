package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Notice struct {
	Id        uint       `json:"id" gorm:"primary_key"`
	Uid       uint32     `json:"uid" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Created   int        `json:"created" gorm:"autoCreateTime"`
	Updated   int        `json:"updated" gorm:"autoUpdateTime"`
	UserId    uint32     `json:"user_id"`
	Content   string     `json:"content"`
}

func (notice *Notice) TableName() string {
	return "notice"
}

func (notice *Notice) BeforeCreate(*gorm.DB) (err error) {
	uid, err := uuid.NewUUID()
	notice.Uid = uid.ID()
	return err
}
