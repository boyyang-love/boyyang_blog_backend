package models

import "time"

type Follow struct {
	Id           uint       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	FollowUserId uint       `json:"follow_user_id"`
	UserId       uint       `json:"user_id"`
}

func (Follow) TableName() string {
	return "follow"
}
