package api

import "github.com/gogf/gf/v2/frame/g"

// 请求get方法
type UserGetInfoReq struct {
	g.Meta `path:"/user/info" method:"get"`
}

type UserGetInfoRes struct {
	Id int `json:"id"`
	// 身份验证key
	IdentityKey string `json:"identity_key"`
	// 载荷
	Payload string `json:"payload"`
}
