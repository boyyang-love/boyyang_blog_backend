package svc

import (
	"blog_server/common/helper"
	"blog_server/internal/config"
	"blog_server/models"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Client *cos.Client
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
		db.AutoMigrate(&models.Follow{})
		fmt.Println("数据库连接成功...")
	}
	client := helper.InitCloudBase(cloudBase.ClientUrl, cloudBase.ClientSecretId, cloudBase.ClientSecretKey)
	if client == nil {
		fmt.Println("对象存储连接失败")
	} else {
		fmt.Println("对象存储连接成功...")
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
		Client: client,
	}
}
