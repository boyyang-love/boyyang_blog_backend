package helper

import (
	"blog_server/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitMysql(args string) (db *gorm.DB, err error) {

	//args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&writeTimeout=%s",
	//		username,
	//		password,
	//		host,
	//		port,
	//		database,
	//		charset,
	//		timeout,
	//	)

	db, err = gorm.Open(mysql.Open(args), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	return db, err
}

func AutoMigrate(db *gorm.DB) (err error) {
	err = db.AutoMigrate(&models.User{})
	err = db.AutoMigrate(&models.Upload{})
	err = db.AutoMigrate(&models.Exhibition{})
	err = db.AutoMigrate(&models.Blog{})
	err = db.AutoMigrate(&models.Comment{})
	err = db.AutoMigrate(&models.Likes{})
	err = db.AutoMigrate(&models.Follow{})
	err = db.AutoMigrate(&models.Tag{})

	return err
}
