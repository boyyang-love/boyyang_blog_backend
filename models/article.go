package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Article struct {
	Id        uint       `json:"id" gorm:"primary_key"`
	Uid       uint32     `json:"uid" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Created   int        `json:"created" gorm:"autoCreateTime"`
	Updated   int        `json:"updated" gorm:"autoUpdateTime"`
	Title     string     `json:"title"`                      // 文章标题
	SubTitle  string     `json:"sub_title"`                  // 文章副标题
	Content   string     `json:"des" gorm:"size:20000"`      // 文章内容
	Cover     string     `json:"cover,omitempty"`            // 背景图片
	UserId    uint32     `json:"user_id"`                    // 文章作者
	Tag       string     `json:"tag,omitempty"`              // 文章标签
	ThumbsUp  *int       `json:"thumbs_up" gorm:"default:0"` // 点赞数
}

func (*Article) TableName() string {
	return "article"
}

func (article *Article) BeforeCreate(*gorm.DB) (err error) {
	uid, err := uuid.NewUUID()
	article.Uid = uid.ID()
	return err
}
