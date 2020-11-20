package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 官方推荐net/http/httptest包进行测试

// 该文件为41. Testing.go的单元测试文件。
// 测试文件的命名规则为为待测试文件添加_test.go后缀。
/*
huangyisan@DESKTOP-59KRN0H MINGW64 ~/Desktop/github/learn-go/learn-gin/Api_Examples/41_Testing (master)
$ go test .
ok      example/Api_Examples/41_Testing 0.255s
 */

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	// 对状态码期望为200，进行断言
	assert.Equal(t, 200, w.Code)
	// 对应答内容为pong，进行断言
	assert.Equal(t, "pong", w.Body.String())
}

