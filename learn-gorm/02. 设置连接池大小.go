package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dsn := "root@tcp(192.168.141.128:3306)/walle?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	// 创建连接池
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	// 设置连接池属性
	// 最大空闲连接数
	sqlDB.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Hour)

	if ok := sqlDB.Close(); ok == nil {
		fmt.Println("数据库关闭成功")
	}

}