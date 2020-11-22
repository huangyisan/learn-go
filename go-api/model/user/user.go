package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	username string
	password string

}

