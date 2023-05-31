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
		fmt.Printf("🐼 Mysql storage initialization failed‼️ 🐼 (%s)", err.Error())
	} else {
		//db.AutoMigrate(&models.User{})
		//db.AutoMigrate(&models.Upload{})
		//db.AutoMigrate(&models.Exhibition{})
		//db.AutoMigrate(&models.Blog{})
		//db.AutoMigrate(&models.Comment{})
		//db.AutoMigrate(&models.Likes{})
		//db.AutoMigrate(&models.Follow{})
		//db.AutoMigrate(&models.Tag{})
		fmt.Println("🐼 Mysql storage initialization successful‼️ 🐼")
	}
	clt := helper.InitCloudBase(cloudBase.ClientUrl, cloudBase.ClientSecretId, cloudBase.ClientSecretKey)
	if clt == nil {
		fmt.Printf("🐼 Object storage initialization failed‼️ 🐼")
		fmt.Println("💫 ")
	} else {
		fmt.Println("🐼 Object storage initialization successful‼️ 🐼")
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
