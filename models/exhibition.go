package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Exhibition struct {
	Id        uint       `json:"id" gorm:"primary_key"`
	Uid       uint32     `json:"uid" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Created   int        `json:"created" gorm:"autoCreateTime"`
	Updated   int        `json:"updated" gorm:"autoUpdateTime"`
	Title     string     `json:"title"`                      // 图片墙标题
	SubTitle  string     `json:"sub_title"`                  // 图片墙副标题
	Des       string     `json:"des"`                        // 图片描述
	Cover     string     `json:"cover"`                      // 图片上传路径
	Tags      *string    `json:"tags"`                       // 图片标签
	ThumbsUp  *int       `json:"thumbs_up" gorm:"default:0"` // 点赞数
	Download  *int       `json:"download" gorm:"default:0"`  // 下载数
	Count     *int       `json:"count" gorm:"default:0"`     // 评论数
	UserId    uint32     `json:"user_id"`                    // 该图片上传者 id
	Status    int        `json:"status" gorm:"default:1"`    // 图片状态 1待审核 2审核通过 3未通过审核 4 公开˚
	RejectRes string     `json:"reject_res"`                 // 状态为3时 驳回原因
	Size      int        `json:"size"`                       // 图片大小
	Px        string     `json:"px"`                         // 图片宽高
	Type      string     `json:"type"`                       // 图片类型
	Rgb       string     `json:"rgb"`                        // 图片主题色
	Palette   string     `json:"palette"`                    // 调色板
}

func (exhibition *Exhibition) TableName() string {
	return "exhibition"
}

func (exhibition *Exhibition) BeforeCreate(*gorm.DB) (err error) {
	uid, err := uuid.NewUUID()
	exhibition.Uid = uid.ID()
	return err
}
