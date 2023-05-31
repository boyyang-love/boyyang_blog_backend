package svc

import (
	"blog_server/common/client"
	"blog_server/common/helper"
	"blog_server/internal/config"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Client *cos.Client
	Hub    *client.Hub
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := c.Mysql
	cloudBase := c.CloudBase
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&writeTimeout=%s",
		mysql.Username,
		mysql.Password,
		mysql.Host,
		mysql.Port,
		mysql.Database,
		mysql.Charset,
		mysql.Timeout,
	)
	db, err := helper.InitMysql(args)
	if err != nil {
		fmt.Printf("数据库连接失败 \n（%s）\n", err.Error())
	} else {
		//db.AutoMigrate(&models.User{})
		//db.AutoMigrate(&models.Upload{})
		//db.AutoMigrate(&models.Exhibition{})
		//db.AutoMigrate(&models.Blog{})
		//db.AutoMigrate(&models.Comment{})
		//db.AutoMigrate(&models.Likes{})
		//db.AutoMigrate(&models.Follow{})
		//db.AutoMigrate(&models.Tag{})
		fmt.Println("数据库连接成功🎇🎇🎇🎇")
	}
	clt := helper.InitCloudBase(cloudBase.ClientUrl, cloudBase.ClientSecretId, cloudBase.ClientSecretKey)
	if clt == nil {
		fmt.Println("对象存储连接失败🧶🧶🧶🧶")
	} else {
		fmt.Println("对象存储连接成功🎇🎇🎇🎇")
	}
	// hub
	hub := client.NewHub()
	return &ServiceContext{
		Config: c,
		DB:     db,
		Client: clt,
		Hub:    hub,
	}
}
