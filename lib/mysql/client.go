package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

type Config struct {
	HostName string
	Username string
	Password string
	Port     int32
	DB       string
}

var DB *gorm.DB

var once sync.Once

func Init(conf *Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username, conf.Password, conf.HostName, conf.Port, conf.DB)
	once.Do(func() {
		client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return
		}
		DB = client
	})
}
