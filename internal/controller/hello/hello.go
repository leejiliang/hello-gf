// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package hello

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"hello-gf/api/hello"
)

type Hello struct {
	msg string
}

func NewHello() *Hello {
	return &Hello{msg: "hello msg"}
}

func (c *Hello) SayHello(req *ghttp.Request) {
	req.Response.Writeln(" Hello Go Frame")
}

func (c *Hello) Params(ctx context.Context, req *hello.ParamsReq) (res *hello.ParamsRes, err error) {
	r := g.RequestFromCtx(ctx)
	//name := r.GetQuery("name", "李四") //李四表示默认值
	//queryMap := r.GetQueryMap()
	//queryMap := r.GetQueryMap(map[string]interface{}{"name": "李四"}) //只获取指定的参数值, 没传name时, 用李四作为默认值
	//queryMap := r.GetQueryMap(g.Map{"name": "李四"}) //map的简写方式, 只获取指定的参数值, 没传name时, 用李四作为默认值
	//username := r.GetForm("username")
	//formMap := r.GetFormMap() // 获取整个表单
	// 转换成结构体测试
	//type person struct {
	//	Username string
	//	Password string
	//}
	//var u person
	//err = r.ParseForm(&u) // 注意这里是pointer, 不要直接传对象.
	//if err == nil {
	//	//r.GetQueryMap()
	//	r.Response.Writeln(u) //拼接之前需要转换
	//
	//}
	//name := r.GetRouter("name")
	//routeMap := r.GetRouterMap()
	//reqMap := r.GetRequestMap()
	reqMap := r.Get("name")
	r.Response.Writeln(reqMap) //拼接之前需要转换
	return
}
