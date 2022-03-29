package functions

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(fileName string) ([]byte, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

// defer 碰到return, 先return然后执行defer.
// 返回值没有命名, 则defer不会影响return的值, 如果返回值有命名,则return的对象defer变更后的值.
func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}
func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
	// fmt.Println(triple(4)) // "12"
}
