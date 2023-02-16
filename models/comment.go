package models

import "time"

type Comment struct {
	Id        uint       `json:"id" gorm:"primary_key" `
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Content   string     `json:"des" gorm:"size:2000"`
	BlogId    uint       `json:"blog_id"`                            // 被评论 博客 id
	UserId    uint       `json:"user_id"`                            // 评论者 id
	UserInfo  User       `json:"user_info" gorm:"foreignKey:UserId"` // 评论者信息
	ThumbsUp  *int       `json:"thumbs_up" gorm:"default:0"`         // 该条评论 点赞数
}

func (Comment) TableName() string {
	return "comment"
}
