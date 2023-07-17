package models

import "time"

type Tag struct {
	Id        uint       `json:"id" gorm:"primary_key" `
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Created   int        `json:"created" gorm:"autoCreateTime"`
	Updated   int        `json:"updated" gorm:"autoUpdateTime"`
	Name      string     `json:"name"`    // 标签名称
	BlogId    uint       `json:"blog_id"` // 博客id
	UserId    uint       `json:"user_id"` // 该标签作者
}

func (Tag) TableName() string {
	return "tag"
}
