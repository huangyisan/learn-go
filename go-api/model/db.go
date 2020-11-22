package model

import (
	"database/sql"
	"fmt"
	"goapi/model/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func InitDB(dbUser,dbPass,dbHost,dbPort,dbname,charset string) *sql.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbname,
		charset,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// 初始化user表
	if ok := db.Migrator().HasTable(&user.User{}); ok {
		log.Println("数据库中已经存在User表")
	}else {
		if err := db.AutoMigrate(&user.User{}); err != nil {
			log.Println("初始化User失败")

		}
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err.Error())
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return sqlDB

}


func isTableExist(db *gorm.DB, table *struct{}) bool {
	if ok := db.Migrator().HasTable(&table); ok {

	}
}

