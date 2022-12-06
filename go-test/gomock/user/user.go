package user

import (
	"fmt"
	"go-test/gomock/person"
)

type User struct {
	Person person.Male
}

func NewUser(p person.Male) *User {
	return &User{Person: p}
}

func (u *User) GetUserInfo(id int64) error {
	return u.Person.Get(id)
}

type aPerson struct {
}

func (a *aPerson) Get(id int64) error {
	fmt.Println(id)
	return nil
}
func k() {
	a := aPerson{}
	NewUser(&a)
}
