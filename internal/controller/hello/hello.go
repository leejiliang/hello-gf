// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package hello

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
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
	//reqMap := r.Get("name")

	r.Response.Writeln(req) //拼接之前需要转换

	return
}

func (c *Hello) Response(ctx context.Context, req *hello.ParamsReq) (res *hello.ParamsRes, err error) {
	//r := g.RequestFromCtx(ctx)
	//r.Response.Writeln(req) //拼接之前需要转换

	//r.Response.Writeln("Hello Go Frame1")
	//r.Response.Write("Hello Go Frame2")

	//r.Response.Writeln("<h1>Hello Go Frame1</h2>")
	//r.Response.Write("<h2>Hello Go Frame2</h2>")
	//r.Response.WriteExit("<h2>Hello Go Frame exit</h2>")
	//r.Response.Writef("<h2>Hello Go Frame3 name is %s , age is %d</h2>", req.Username, req.Age)

	//r.Response.WriteJson(req) // 返回信息: Content-Type: application/json,
	res = &hello.ParamsRes{
		Username: "winston",
		Password: "123456a",
		Age:      122,
	}
	//r.Response.Write(req) // Content-Type: text/plain; charset=utf-8
	err = gerror.New("服务器GG")
	return
}

func (c *Hello) Db(req *ghttp.Request) {
	//req.Response.Writeln(g.Model("book"))
	model := g.Model("book")
	//record, err := model.One()
	//record, err := model.Fields("id,name").All()
	//record, err := model.Fields("id", "name").All()
	//record, err := model.FieldsEx("id", "name").All()
	//record, err := model.Value("name") // 表示一条数据的一个具体字段值
	//record, err := model.Array("name") // 表示查询某一列数据的全部值
	//min, err := g.Model("book").Min("price")
	//max, err := g.Model("book").Max("price")
	//avg, err := g.Model("book").Avg("price")
	//count, err := g.Model("book").Count()
	//all, err := model.Where("id", 3).All()
	//all, err := model.Where("id<", 3).All()
	//all, err := model.Where("id<?", 2).All()
	//all, err := model.WhereLTE("id", 4).WhereGTE("id", 2).All()
	//all, err := model.WhereIn("id", g.Array{3, 4, 5}).All()
	//all, err := model.WhereIn("id", g.Array{3, 4, 5}).WhereOr("id", 1).All()
	all, err := model.Order("id", "DESC").OrderAsc("price").Group("name").All()
	if err == nil {
		req.Response.WriteJson(all)
		//req.Response.Writeln(record["name"])
		//req.Response.Writeln(g.Map{
		//	"min":   min,
		//	"max":   max,
		//	"avg":   avg,
		//	"count": count,
		//})
		//req.Response.Writeln(record[0]["name"])
	}
}
