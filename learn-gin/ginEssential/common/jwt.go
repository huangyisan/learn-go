package common

import (
	"ginEssential/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtkey = []byte("a_secret_token")

type Claims struct {
	UserId uint
	// 继承jwt的标准Claims
	jwt.StandardClaims
}

// 发放token
func ReleaseToken(user model.User) (string, error) {
	// 过期时间7天
	expirationTime := time.Now().Add(7* 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer: "ys",
			Subject: "user token",
		},
	}

	// 指定算法hs256 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用jwtkey签名
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}