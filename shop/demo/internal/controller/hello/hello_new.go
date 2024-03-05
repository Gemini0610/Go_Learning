// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package hello

import (
	"demo/api/hello"
)
// 承上启下的作用
// 承接上面api传来的参数
// 启下：调用后面service层和logic层的业务逻辑
type ControllerV1 struct{}

func NewV1() hello.IHelloV1 {
	return &ControllerV1{}
}

