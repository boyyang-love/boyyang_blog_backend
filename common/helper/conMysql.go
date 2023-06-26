package helper

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	db, err = gorm.Open(mysql.Open(args), &gorm.Config{})

	return db, err
}
