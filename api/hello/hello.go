// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package hello

import "github.com/gogf/gf/v2/frame/g"

type ParamsReq struct {
	//g.Meta `method: "all" `
	g.Meta `path:"/params" method:"all" ` //路径参数
	Username string
	Password string
	Age int
}

type ParamsRes struct {}