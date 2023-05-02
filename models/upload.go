package models

import "time"

type Upload struct {
	Id        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Created   int        `json:"created" gorm:"autoCreateTime"`
	Updated   int        `json:"updated" gorm:"autoUpdateTime"`
	Hash      string     `json:"hash"`
	FileName  string     `json:"file_name"`
	FilePath  string     `json:"file_path"`
	Ext       string     `json:"ext"`
	Size      int64      `json:"size"`
	UserId    uint       `json:"user_id"`
}

func (Upload) TableName() string {
	return "upload"
}
