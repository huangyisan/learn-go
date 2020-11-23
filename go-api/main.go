package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goapi/router"
	"os"
	"goapi/model"
)

func initConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func main() {
	// 读取yaml配置
	initConfig()

	// db初始化
	db := model.InitDB()
	
	defer db.Close()

	// 创建gin实例
	r := gin.Default()

	router.Router(r)

	if err := r.Run(":7777"); err != nil {
		panic(err.Error())
	}
}