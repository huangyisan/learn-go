package user

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var User1 = User{
	ID:       1,
	Username: "username",
	Password: "password"}
