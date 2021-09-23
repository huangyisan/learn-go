package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20); not null"`
	Telephone string `gorm:"type:varchar(11); not null;unique"`
	Password string `gorm:"size:255; not null"`
}

func main() {
	db := initDB()

	r := gin.Default()
	r.POST("/api/user/register", func(c *gin.Context) {

		// 获取参数
		name := c.PostForm("name")
		password := c.PostForm("password")
		telephone := c.PostForm("telephone")

		// 参数校验
		if len(telephone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "手机号码必须为11位",
			})
		}

		if len(password) <6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "密码长度必须大于6位",
			})
		}

		if len(name) == 0 {
			name = randomName()
		}

		// 判断手机号码是否存在
		if isTelephoneExist(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "手机已经被注册",
			})
		}

		// 注册用户
		newUser := User{
			Name: name,
			Telephone: telephone,
			Password: password,
		}

		db.Create(&newUser)

		c.JSON(200, gin.H{
			"message": "注册成功",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func randomName() string {
	//var letters = []string{"abc","cc"}
	//fmt.Println(letters)

	var letters = []byte("abccsasdfwerqxcsaer")
	result := make([]byte, 8)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
		//}

	}
	return string(result)
}

func initDB() *gorm.DB {
	user := "root"
	pass := "huangyisan"
	ip := "192.168.200.128"
	port := "3306"
	dbname := "ginessential"
	charset := "utf8mb4"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",user, pass, ip, port, dbname, charset )

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to db %s", ip))
	}

	// 自动创建user表
	db.AutoMigrate(&User{})

	return db
}