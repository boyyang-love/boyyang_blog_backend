package models

import "time"

type Likes struct {
	Id           uint       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	ExhibitionId uint       `json:"exhibition_id"`
	UserId       uint       `json:"user_id"`
}

func (Likes) TableName() string {
	return "likes"
}
