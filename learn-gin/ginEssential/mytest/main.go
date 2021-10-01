package main

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20); not null"`
	Telephone string `gorm:"type:varchar(11); not null;unique"`
	Password  string `gorm:"size:255; not null"`
}

func main() {
	var user interface{}
	user = User{
		Name:      "yisan",
		Telephone: "123",
		Password:  "123123",
	}
	// 断言
	a := user.(User)
	fmt.Println(a)
}
