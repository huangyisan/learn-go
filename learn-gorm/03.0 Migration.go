package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID			 uint
	Name         string
	Age          uint8
}

func main() {
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dsn := "root@tcp(192.168.141.128:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
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

	// migrate数据， 建立user表
	db.AutoMigrate(&User{})

	// 查看是否建立
	fmt.Println("是否存在?", db.Migrator().HasTable(&User{}))



	if ok := sqlDB.Close(); ok == nil {
		fmt.Println("数据库关闭成功")
	}

}