package main

import (
	"github.com/go-playground/validator/v10"
	"time"
)

// 当默认的validator无法满足的时候，需要自定义方法写验证函数
// 自定义验证函数跟随在tag的"bind"后面
// 在main函数中注册自定义的验证函数


func main() {
	// 定义book的结构体
	type Booking struct {
		// 属性首字母大写， 指定time.Time类型， binding内bookabledate为自定义校验函数，time_format表示传入的时间格式样式
		CheckIn time.Time `json:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
		// gtfield表示比指定字段要大，这边业绩是checkout字段要比checkin来的大
		CheckOut time.Time `json:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
	}
	// 自定义bookabledate验证函数, validator.Func类型为对传入的值进行验证，如果通过则返回true
	var bookabledate validator.Func = func(fl validator.FieldLevel) bool {
		// Field()为获取原先的方法， 然后使用反射的方式将其转换为time.Time类型
		// 反射的使用可以参考previous/进阶部分/反射/1.快速案例-对基本数据类型反射.go
		date, ok := fl.Field().Interface().(time.Time)
			if ok {
				today := time.Now()
				// After函数用来比较today对象是否比date对象来的晚
				if today.After(date) {
					return false
				}
			}
			return true
	}
}

func main() {
	
}
