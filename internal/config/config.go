package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Mysql struct {
		Host     string
		Port     int
		Database string
		Username string
		Password string
		Charset  string
		Timeout  string
	}
	CloudBase struct {
		ClientUrl       string
		ClientSecretId  string
		ClientSecretKey string
	}
}
