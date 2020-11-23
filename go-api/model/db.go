package model

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"goapi/model/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// 单实例模式， 当InitDB实例初始化后，DB就有对象了
var DB *gorm.DB

func InitDB() *sql.DB {

	dbUser := viper.GetString("Database.dbUser")
	dbPass := viper.GetString("Database.dbPass")
	dbHost := viper.GetString("Database.dbHost")
	dbPort := viper.GetString("Database.dbPort")
	dbname := viper.GetString("Database.dbname")
	charset := viper.GetString("Database.charset")

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

	// 单实例赋予对象
	DB = db

	return sqlDB
}

func GetDB() *gorm.DB {
	return DB
}

func migrateTable(db *gorm.DB, table interface{}) {
	if ok := isTableExist(db, table); !ok {
		if err := db.AutoMigrate(table); err != nil {
			log.Printf("初始化%+v失败",table)
		}else {
			log.Printf("初始化%+v成功",table)
		}
	}else {
		log.Printf("数据库已经存在%+v表",table)
	}


}

func isTableExist(db *gorm.DB, table interface{}) bool {
	if ok := db.Migrator().HasTable(table); ok {
		fmt.Println("true")
		// 已经存在该表
		return true
	}
	// 不存在该表
	fmt.Println("false")
	return false
}

