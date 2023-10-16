package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	Id        uint       `json:"id" gorm:"primary_key"`
	Uid       uint32     `json:"uid" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Created   int        `json:"created" gorm:"autoCreateTime"`
	Updated   int        `json:"updated" gorm:"autoUpdateTime"`
	Type      string     `json:"type"`                       // image blog article
	Content   string     `json:"des" gorm:"size:2000"`       // 评论内容
	ContentId uint32     `json:"content_id"`                 // 被评论 id
	UserId    uint32     `json:"user_id"`                    // 评论者 id
	ThumbsUp  *int       `json:"thumbs_up" gorm:"default:0"` // 该条评论 点赞数
}

func (comment *Comment) TableName() string {
	return "comment"
}

func (comment *Comment) BeforeCreate(*gorm.DB) (err error) {
	uid, err := uuid.NewUUID()
	comment.Uid = uid.ID()
	return err
}
