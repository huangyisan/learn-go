package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:265"`
	Password string `gorm:"size:265"`

}

