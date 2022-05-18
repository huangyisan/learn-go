package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	UserID uint64 `json:"user_id"`
	Title  string `json:"title"`
}

// 执行todo之前进行token校验,看是否合法并且可用.
func CreateTodo(c *gin.Context) {
	var userId uint64
	var td *Todo
	if err := c.ShouldBindJSON(&td); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}
	tokenAuth, err := ExtractTokenMetadata(c.Request)
	fmt.Printf("%v", tokenAuth)
	if err != nil {

		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	userId, err = FetchAuth(tokenAuth)
	fmt.Printf("%v", userId)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	td.UserID = userId

	//you can proceed to save the Todo to a database
	//but we will just return it to the caller here:
	c.JSON(http.StatusCreated, td)
}
