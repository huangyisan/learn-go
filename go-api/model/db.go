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
	migrateTable(db, &user.User{})

	// 初始化db pool  设置pool 属性
	sqlDB, err := db.DB()
	if err != nil {
		panic(err.Error())
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return sqlDB
}

func migrateTable(db *gorm.DB, table interface{}) {
	if ok := isTableExist(db, table); !ok {
		if err := db.AutoMigrate(&table); err != nil {
			log.Printf("初始化%v失败",table)
		}
		log.Printf("初始化%v成功",table)
	}
	log.Printf("数据库已经存在%v表",table)

}

func isTableExist(db *gorm.DB, table interface{}) bool {
	if ok := db.Migrator().HasTable(&table); ok {
		// 已经存在该表
		return true
	}
	// 不存在该表
	return false
}

