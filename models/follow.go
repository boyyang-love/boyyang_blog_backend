package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Follow struct {
	Id           uint       `json:"id" gorm:"primary_key"`
	Uid          uint32     `json:"uid" gorm:"primary_key"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	Created      int        `json:"created" gorm:"autoCreateTime"`
	Updated      int        `json:"updated" gorm:"autoUpdateTime"`
	FollowUserId uint32     `json:"follow_user_id"` // 被关注者 id
	UserId       uint32     `json:"user_id"`        // 关注者 id
	FollowType   bool       `json:"follow_type"`    // true 关注 false 取消关注
}

func (*Follow) TableName() string {
	return "follow"
}

func (follow *Follow) BeforeCreate(*gorm.DB) (err error) {
	uid, err := uuid.NewUUID()
	follow.Uid = uid.ID()
	return err
}
