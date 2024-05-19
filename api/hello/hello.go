// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package hello

import "github.com/gogf/gf/v2/frame/g"

type ParamsReq struct {
	//g.Meta `method: "all" `
	g.Meta `path:"/params/{name}" method:"all" ` //路径参数
}

type ParamsRes struct {}