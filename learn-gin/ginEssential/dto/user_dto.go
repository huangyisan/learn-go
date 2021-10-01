package dto

import "ginEssential/model"

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User) UserDto {
	// 只需要name telephone信息,不需要password
	return UserDto{
		user.Name,
		user.Telephone,
	}
}
