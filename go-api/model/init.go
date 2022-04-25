package model

import "gorm.io/gorm"

// 因为一个 API 服务器可能需要同时访问多个数据库，为了对多个数据库进行初始化和连接管理，这里定义了一个叫 Database 的 struct
type DataBase struct {
	Self *gorm.DB
}

func (db *DataBase) Init() {
	DB = &DataBase{
		viper.G
	}
}
