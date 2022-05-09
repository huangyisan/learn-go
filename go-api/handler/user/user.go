package user

import "go-api/model"

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

// 使用偏移量对用户进行分页输出
type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount int64             `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}
