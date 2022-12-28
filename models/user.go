package models

import "time"

type User struct {
	Id        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Username  string     `form:"username" json:"username"`
	Password  string     `form:"password" json:"password"`
	Gender    int        `form:"gender" json:"gender"`
	Age       int        `form:"age" json:"age"`
	Birthday  *int64     `form:"birthday" json:"birthday" gorm:"default:0"`
	Address   string     `form:"address" json:"address"`
	Tel       *int       `form:"tel" json:"tel"`
	Email     *string    `form:"email" json:"email" gorm:"default:xxxxxx@qq.com"`
	AvatarUrl string     `form:"avatar_url" json:"avatar_url"`
}

func (User) TableName() string {
	return "user"
}
