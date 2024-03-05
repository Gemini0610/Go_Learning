package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RotationReq struct {
	g.Meta `path:"/backend/rotation/add" tags:"Rotation" method:"post" summary:"You first rotation api"`
	PicUrl string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort   int    `json:"sort"    dc:"排序"`
}

type RotationRes struct {
	//todo
	//g.Meta `mime:"text/html" example:"string"`
	RotationId int `json:"rotationId"`
}
type RotationDeleteReq struct {
	//path:对应的接口url
	g.Meta `path:"/backend/rotation/delete" method:"delete" tags:"轮播图" summary:"删除轮播图接口"`
	Id     uint `v:"min:1#请选择需要删除的轮播图" dc:"轮播图id"`
}
type RotationDeleteRes struct{}

type RotationUpdateReq struct {
	g.Meta `path:"/rotation/update/{Id}" method:"post" tags:"轮播图" summary:"修改轮播图接口"`
	Id     uint   `json:"id"      v:"min:1#请选择需要修改的轮播图" dc:"轮播图Id"`
	PicUrl string `json:"pic_url"    v:"required#轮播图图片链接不能为空" dc:"轮播图图片链接"`
	Link   string `json:"link"    v:"required#轮播图跳转链接不能为空" dc:"轮播图跳转链接"`
	Sort   int    `json:"sort"     dc:"轮播图排序"`
}
type RotationUpdateRes struct {
	Id uint `json:"id"`
}

// 取值GET
type RotationGetListCommonReq struct {
	//需要定义path
	g.Meta `path:"/rotation/list" method:"get" tags:"轮播图" summary:"轮播图列表接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型`
	CommonPaginationReq
}

// 输出
type RotationGetListCommonRes struct {
	//todo
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`
	//数据侧与业务侧分离，所以是interface{}
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}