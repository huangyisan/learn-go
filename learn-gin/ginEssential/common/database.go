package common

import (
	"fmt"
	"ginEssential/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	user := "root"
	pass := "huangyisan"
	// ip := "192.168.200.128"
	ip := "47.102.119.67"
	port := "3306"
	dbname := "ginessential"
	charset := "utf8mb4"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",user, pass, ip, port, dbname, charset )

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to db %s", ip))
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