package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id              uint       `gorm:"primary_key" json:"id"`
	Uid             uint       `json:"uid" gorm:"primary_key"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	Created         int        `json:"created" gorm:"autoCreateTime"`
	Updated         int        `json:"updated" gorm:"autoUpdateTime"`
	Username        string     `form:"username" json:"username"`
	Password        string     `form:"password" json:"password"`
	Gender          int        `form:"gender" json:"gender"`
	Age             int        `form:"age" json:"age"`
	Birthday        int64      `form:"birthday" json:"birthday" gorm:"default:0"`
	Address         string     `form:"address" json:"address"`
	Tel             int        `form:"tel" json:"tel"`
	Email           string     `form:"email" json:"email"`
	Qq              int        `form:"qq" json:"qq"`
	Wechat          string     `form:"wechat" json:"wechat"`
	GitHub          string     `form:"git_hub" json:"git_hub"`
	AvatarUrl       string     `form:"avatar_url" json:"avatar_url"`
	BackgroundImage string     `form:"background_image" json:"background_image"`
	Motto           string     `form:"motto" json:"motto"`
}

func (user *User) TableName() string {
	return "user"
}

func (user *User) BeforeCreate(db *gorm.DB) (err error) {
	user.Uid = uint(uuid.New().ID())
	return
}
