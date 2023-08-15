package svc

import (
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
		fmt.Printf("🐼 Mysql database initialization failed‼️ 🐼 (%s)", err.Error())
	} else {
		err := helper.AutoMigrate(db)
		if err != nil {
			fmt.Println("🐼 Mysql database automigrage failed‼️ 🐼")
		}
		fmt.Println("🐼 Mysql database initialization successful‼️ 🐼")
	}
	clt := helper.InitCloudBase(cloudBase.ClientUrl, cloudBase.ClientSecretId, cloudBase.ClientSecretKey)
	if clt == nil {
		fmt.Printf("🐼 Object storage initialization failed‼️ 🐼")
	} else {
		fmt.Println("🐼 Object storage initialization successful‼️ 🐼")
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
		Client: clt,
	}
}
