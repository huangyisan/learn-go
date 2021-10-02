package common

import (
	"fmt"
	"ginEssential/model"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	user := viper.GetString("datasource.dbUser")
	password := viper.GetString("datasource.password")
	// ip := "192.168.200.128"
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	dbname := viper.GetString("datasource.dbName")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", user, password, host, port, dbname, charset)

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to db %s", host))
	}

	// 自动创建user表
	db.AutoMigrate(&model.User{})

	// DB单例模式
	DB = db

	return db
}

// 定义方法获取db实例
func GetDB() *gorm.DB {
	return DB
}
