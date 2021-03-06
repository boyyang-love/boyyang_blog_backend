/**
 * @Author: boyyang
 * @Date: 2022-04-03 00:02:39
 * @LastEditTime: 2022-07-15 13:58:03
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\models\pictureWalls.go
 * [如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package models

import (
	"time"
)

type PictureWall struct {
	// gorm.Model
	ID           uint           `gorm:"primary_key" json:"id" form:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    *time.Time     `sql:"index" json:"deleted_at"`
	Url          string         `json:"url" form:"url"`                                                    //图片地址
	FileName     string         `json:"file_name" form:"file_name"`                                        //图片名称
	Name         string         `json:"name" form:"name"`                                                  //名称
	Des          string         `json:"des" form:"des"`                                                    //图片描述
	Type         *int           `json:"type" form:"type" gorm:"default:0"`                                 //图片类型 0 pc端 1 手机端
	Hidden       *int           `json:"hidden" form:"hidden" gorm:"default:0"`                             //是否隐藏 0 否 1是
	Status       *int           `json:"status" form:"status" gorm:"default:0"`                             //状态 0 待审核 1 审核通过 2 审核不通过
	ThumbsUp     int            `json:"thumbs_up" form:"thumbs_up"`                                        //点赞数
	ThumbsUpList []ThumbsUp     `json:"thumbs_up_list" form:"thumbs_up_list" gorm:"foreignKey:ThumbsUpId"` //点赞列表
	Tags         []ImagesTag    `json:"tags" form:"tags"`                                                  //标签类型
	UserID       int            `json:"user_id" form:"user_id"`                                            //用户id
	Author       User           `json:"author" gorm:"foreignKey:UserID"`                                   //gorm:"foreignKey:UserID"
	LeaveMessage []LeaveMessage `json:"leave_message" form:"leave_message" gorm:"foreignKey:CommentId"`    //留言列表
}

func (PictureWall) TableName() string {
	return "pictureWalls"
}
