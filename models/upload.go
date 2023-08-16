package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Upload struct {
	Id        uint       `json:"id" gorm:"primary_key"`
	Uid       uint       `json:"uid" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Created   int        `json:"created" gorm:"autoCreateTime"`
	Updated   int        `json:"updated" gorm:"autoUpdateTime"`
	Hash      string     `json:"hash"`      // 文件哈希
	FileName  string     `json:"file_name"` // 文件名称
	FilePath  string     `json:"file_path"` // 文件路径
	Ext       string     `json:"ext"`       // 文件后缀
	Size      int64      `json:"size"`      // 文件大小
	UserId    uint       `json:"user_id"`   // 用户id
}

func (*Upload) TableName() string {
	return "upload"
}

func (upload *Upload) BeforeCreate(db *gorm.DB) (err error) {
	uid, err := uuid.NewUUID()
	upload.Uid = uint(uid.ID())
	return err
}
