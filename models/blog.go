package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Blog struct {
	Id           uint       `json:"id" gorm:"primary_key"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	Created      int        `json:"created" gorm:"autoCreateTime"`
	Updated      int        `json:"updated" gorm:"autoUpdateTime"`
	Title        string     `json:"title"`                              // 博客标题
	SubTitle     string     `json:"sub_title"`                          // 博客副标题
	Content      string     `json:"des" gorm:"size:10000"`              // 博客内容
	Cover        string     `json:"cover,omitempty"`                    // 背景图片
	UserId       uint       `json:"user_id"`                            // 博客作者
	UserInfo     User       `json:"user_info" gorm:"foreignKey:UserId"` // 用户信息
	Tag          string     `json:"tag,omitempty"`                      // 博客标签
	ThumbsUp     *int       `json:"thumbs_up" gorm:"default:0"`         // 点赞数
	ThumbsUpList string     `json:"thumbs_up_list"`                     // 点赞id集合
	Collection   *int       `json:"collection" gorm:"default:0"`        // 收藏数
}

func (*Blog) TableName() string {
	return "blog"
}

func (blog *Blog) BeforeCreate(db *gorm.DB) (err error) {
	blog.Id = uint(uuid.New().ID())
	return
}
