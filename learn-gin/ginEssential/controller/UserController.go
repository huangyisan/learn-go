package controller

import (
	"fmt"
	"ginEssential/common"
	"ginEssential/dto"
	"ginEssential/model"
	"ginEssential/response"
	"ginEssential/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {

	DB := common.GetDB()

	// 获取参数
	name := c.PostForm("name")
	password := c.PostForm("password")
	telephone := c.PostForm("telephone")

	// 参数校验
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号码必须为11位")
		return
	}

	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码长度必须大于6位")
		return
	}

	// 判断手机号码是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机已经被注册")
		return
	}

	if len(name) == 0 {
		name = util.RandomName()
	}

	// 注册用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 422, nil, "加密错误")
		return
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
	}

	DB.Create(&newUser)

	response.Success(c, nil, "注册成功")
}

func UserLogin(c *gin.Context) {
	DB := common.GetDB()
	// args
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	// validate
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号码必须为11位")
		return
	}

	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码长度必须大于6位")
		return
	}

	// 手机号码验证
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	// 判断密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusUnauthorized, 403, nil, "密码错误")
		return
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 422, nil, "系统异常")
		return
	}

	response.Success(c, gin.H{"token": token}, "登陆成功")
	// 返回结果
}

// 查询信息
func Info(c *gin.Context) {
	// 用户此时为登陆状态，则从上下文中获取用户信息
	user, _ := c.Get("user")
	// a := user.(model.User)
	response.Success(c, gin.H{
		// Get方法返回的类型为interface, 这里user成了interface类型,则需要断言转换为model.User
		"user": dto.ToUserDto(user.(model.User)),
	}, "")

}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	fmt.Println(user.ID)
	return user.ID != 0
}
