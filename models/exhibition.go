package models

import "time"

type Exhibition struct {
	Id        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Created   int        `json:"created" gorm:"autoCreateTime"`
	Updated   int        `json:"updated" gorm:"autoUpdateTime"`
	Title     string     `json:"title"`
	SubTitle  string     `json:"sub_title"`
	Des       string     `json:"des"`
	Cover     string     `json:"cover"`                      // 图片上传路径
	ThumbsUp  *int       `json:"thumbs_up" gorm:"default:0"` // 点赞数
	UserId    uint       `json:"user_id"`                    // 该图片上传者 id
	Status    int        `json:"status" gorm:"default:1"`    // 1待审核 2审核通过 3未通过审核
	RejectRes string     `json:"reject_res"`                 // 驳回原因
}

func (Exhibition) TableName() string {
	return "exhibition"
}
