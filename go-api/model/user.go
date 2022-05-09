package model

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"go-api/pkg/auth"
	"go-api/pkg/constvar"
)

// User represents a registered user.
type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"` // 添加validate tag使用validate包进行验证
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

// Validate the fields.
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// 密码加密
func (u *UserModel) Encrypt() (err error) {
	// 这个地方对Password进行了加密
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// 密码和加密后的密码比较
func (u *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// 创建用户
func (u *UserModel) Create() (err error) {
	return DB.Self.Create(&u).Error
}

// 删除用户 用id删除
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.Id = id
	return DB.Self.Delete(&user).Error
}

// Update updates an user account information.
func (u *UserModel) Update() error {
	return DB.Self.Save(u).Error
}

// 获取用户信息
func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("username = ?", username).First(&u)
	return u, d.Error
}

// ListUser List all users
func ListUser(username string, offset, limit int) ([]*UserModel, int64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit

	}
	users := make([]*UserModel, 0)
	var count int64
	// 模糊查询 https://blog.csdn.net/qq_35167735/article/details/107862234
	where := fmt.Sprintf("username like '%%%s%%'", username)
	// 查询符合条件的个数
	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err

	}

	// 查询具体的user信息
	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}
	return users, count, nil
}
