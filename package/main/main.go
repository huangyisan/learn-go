// 如果要编译成可执行文件，则将这个包申明为main包。
package main
//包路径会从 $GOPATH下的src开始寻找。
// 也可以使用相对路径
import (
	"fmt"
	// 在引入前面，可以写别名
	util "../utils"
)


func main() {
	//fmt.Println(utils.Cal(1,2,'+'))
	fmt.Println(util.Cal(1,2,'+'))
	fmt.Println(util.Cal2(1,2,'+'))
}