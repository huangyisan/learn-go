package middleware

import (
	"ginEssential/common"
	"ginEssential/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// gin中间件，一个函数返回gin.HandlerFunc
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		// validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized,
				gin.H{"message": "权限不足"}, )
			// 抛弃该次请求
			ctx.Abort()
			return
		}

		// 提取token有效部分
		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "权限不足",
			})
			ctx.Abort()
			return
		}

		// 通过验证, 获取claims中的userid
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 用户是否存在，不存在这token无效
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized,
				gin.H{"message": "权限不足"},
				)
			ctx.Abort()
			return
		}

		// 用户存在，将user信息写入上下文
		ctx.Set("user", user)

		// 继续执行下面的行为。Next()方法只能存在middleware中
		ctx.Next()

	}
}