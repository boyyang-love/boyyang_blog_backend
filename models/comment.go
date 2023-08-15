package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	Id        uint       `json:"id" gorm:"primary_key" `
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Created   int        `json:"created" gorm:"autoCreateTime"`
	Updated   int        `json:"updated" gorm:"autoUpdateTime"`
	Content   string     `json:"des" gorm:"size:2000"`       // 评论内容
	BlogId    uint       `json:"blog_id"`                    // 被评论 博客 id
	UserId    uint       `json:"user_id"`                    // 评论者 id
	ThumbsUp  *int       `json:"thumbs_up" gorm:"default:0"` // 该条评论 点赞数
}

func (comment *Comment) TableName() string {
	return "comment"
}

func (comment *Comment) BeforeCreate(db *gorm.DB) (err error) {
	comment.Id = uint(uuid.New().ID())
	return
}
