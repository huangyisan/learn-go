package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dsn := "root@tcp(192.168.141.128:3306)/walle?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	if ok := sqlDB.Close(); ok == nil {
		fmt.Println("数据库关闭成功")
	}

}