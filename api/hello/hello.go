// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package hello

import "github.com/gogf/gf/v2/frame/g"

type ParamsReq struct {
	//g.Meta `method: "all" `
	g.Meta `path:"/params" method:"all" ` //路径参数
	Username string `p:"name" d:"admin"` // p: 表示param: 参数别名, 请求传参可以使用的属性名, d表示默认值.
	Password string
	Age int
}

type ParamsRes struct {}