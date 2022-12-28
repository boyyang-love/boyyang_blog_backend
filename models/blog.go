package models

import "time"

type Blog struct {
	Id        uint       `json:"id" gorm:"primary_key" `
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Title     string     `json:"title"`
	SubTitle  string     `json:"sub_title"`
	Content   string     `json:"des" gorm:"size:2000"`
	Cover     string     `json:"cover,omitempty"`
	UserId    uint       `json:"user_id"`
	UserInfo  User       `json:"user_info" gorm:"foreignKey:UserId"`
	Tag       string     `json:"tag,omitempty"`
	ThumbsUp  *int       `json:"thumbs_up" gorm:"default:0"`
	Comments  []Comment  `json:"comments" gorm:"foreignKey:BlogId"`
}

func (Blog) TableName() string {
	return "blog"
}
