package model

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zxmrlc/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 因为一个 API 服务器可能需要同时访问多个数据库，为了对多个数据库进行初始化和连接管理，这里定义了一个叫 Database 的 struct
type DataBase struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var DB *DataBase

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local",
	)
	db, err := gorm.Open(mysql.Open(config), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}
	// set for db connection
	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Errorf(err, "Database setup failed")
	}
	sqlDB.SetMaxIdleConns(0)
}

// used for cli
func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
	)
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

// docker db
func InitDockerDB() *gorm.DB {
	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"))
}

func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}

func (db *DataBase) Init() {
	DB = &DataBase{
		Self:   GetSelfDB(),
		Docker: GetDockerDB(),
	}
}

func (db *DataBase) Close() {
	sqlDBCli, _ := db.Self.DB()
	sqlDBCli.Close()
	sqlDBDocker, _ := db.Docker.DB()
	sqlDBDocker.Close()
}
