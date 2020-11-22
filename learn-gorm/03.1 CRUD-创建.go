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

	// 创建数据
	user := User{Name: "Jinzhu", Age: 18}
	// 通过指针来创建
	result := db.Create(&user)
	fmt.Println(user.ID)
	fmt.Println(user.Name)
	fmt.Println(user.Age)
	fmt.Println("插入的error为", result.Error)
	fmt.Println("插入的条数为", result.RowsAffected )


	if ok := sqlDB.Close(); ok == nil {
	}

}