// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package hello

import "github.com/gogf/gf/v2/frame/g"

type ParamsReq struct {
	g.Meta `method: "all" `
	//g.Meta `path:"/params" method:"all" ` //路径参数
	Username string `p:"name" d:"admin"` // p: 表示param: 参数别名, 请求传参可以使用的属性名, d表示默认值.
	Password string
	Age int
}

type ParamsRes struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age int `json:"age"`
}

type ValidReq struct {
	Username string `p:"username" v:"required"`
	Password string `p:"password" v:"required"`
	Age int `p:"age" v:"required|integer|min:0|max:180#Ageb必填|年龄必须是整型|最小值为0|最大值为180"`
}

type ValidRes struct {}