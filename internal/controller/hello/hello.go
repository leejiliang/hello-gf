// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package hello

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"hello-gf/api/hello"
	"hello-gf/internal/dao"
	"hello-gf/internal/model/do"
	"hello-gf/internal/model/entity"
	"hello-gf/internal/service"
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
	//tx, err := g.DB().Begin(req.Context())
	//if err != nil {
	//	req.Response.WriteExit("事务开启失败")
	//}
	////model := g.Model("book")
	//model := tx.Model("book")
	//result, err := model.Delete("id>?", 9)

	//err := g.DB().Transaction(req.Context(), func(ctx context.Context, tx gdb.TX) error {
	//	model := tx.Model("book")
	//	result, err := model.Delete("id>", 5)
	//	if err == nil {
	//		req.Response.Write(result)
	//	} else {
	//		req.Response.Write(err)
	//	}
	//	return err
	//})
	//if err != nil {
	//	return
	//}
	// 执行原生sql语句
	//db := g.DB()
	//db.Query()
	//db.Exec()

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
	//all, err := model.Order("id", "DESC").OrderAsc("price").Group("name").All()
	//all, err := model.Limit(3, 2).All()
	//var data = g.Map{
	//	"price":        450.2,
	//	"publish_time": "2024-05-20",
	//}
	////result, err := model.Where(g.Map{
	////	"author": "无名大侠",
	////}).Data(data).Update()
	//result, err := model.Where(g.Map{
	//	"author": "无名大侠",
	//}).Update(data)
	type Book struct {
		Id          int
		BookName    string `orm:"name"` //orm 标签用于映射类型属性和数据库字段名
		Author      string `json:"author"`
		Price       float32
		PublishTime *gtime.Time
	}

	//data := g.Map{
	//	//"id":           9,
	//	"name":         "go入门实战9",
	//	"author":       "无名大侠",
	//	"price":        200,
	//	"publish_time": "2024-01-03",
	//}
	//data := Book{
	//	//"id":           9,
	//	BookName: "go入门实战aa",
	//	Author:   "无名大侠",
	//	Price:    1.0,
	//	//PublishTime:     gtime.Now(),
	//	PublishTime: gtime.New("2023-09-09"),
	//}
	//var books = g.Array{
	//	Book{
	//		BookName: "go入门实战aa1",
	//		Author:   "无名大侠",
	//		Price:    1.0,
	//		//PublishTime:     gtime.Now(),
	//		PublishTime: gtime.New("2023-09-09"),
	//	},
	//	Book{
	//		BookName: "go入门实战aa2",
	//		Author:   "无名大侠",
	//		Price:    1.0,
	//		//PublishTime:     gtime.Now(),
	//		PublishTime: gtime.New("2023-09-10"),
	//	},
	//}
	//result, err := model.Data(data).Insert()
	//result, err := model.Replace(data)
	//result, err := model.Insert(data)
	//result, err := model.Save(books)
	//result, err := model.InsertAndGetId(data)

	//var book *Book // 指针类型, 查询不到数据时， 指针为nil, 没有具体的err信息
	//var book []Book // 指针类型, 查询不到数据时， 指针为nil, 没有具体的err信息
	//var book Book // 非指针类型, 查询不到数据时， error会为： sql: no rows in result set
	//all, err := model.Page(1, 2).All()
	//err := model.Scan(&book)
	//result, err := model.Delete("id>?", 10)
	//result, err := model.Where("id>?", 9).All()

	//if err == nil {
	//	req.Response.WriteJson(result)
	//	tx.Commit()
	//	//req.Response.Writeln(record["name"])
	//	//req.Response.Writeln(g.Map{
	//	//	"min":   min,
	//	//	"max":   max,
	//	//	"avg":   avg,
	//	//	"count": count,
	//	//})
	//	//req.Response.Writeln(record[0]["name"])
	//} else {
	//	tx.Rollback()
	//	req.Response.Write("错误信息:" + err.Error())
	//}
	ctx := dao.Book.Ctx(req.Context())
	//ctx = ctx.WhereGTE("id", 1)
	//ctx = ctx.WhereLTE("id", 3)
	//all, err := ctx.All()
	//if err == nil {
	//	req.Response.Writeln(all)
	//} else {
	//	req.Response.Write(err)
	//}
	//book := entity.Book{Author: "猫腻", Price: 20000} // 其他属性如果不给值，会使用默认值， 如果更新数据库会导致数据库
	book := do.Book{Price: 300}
	//其他字段被更新为默认值。
	//result, err := ctx.Where("id", 1).Fields("author,price").Data(book).Update() // 只更新部分字段
	result, err := ctx.Where(do.Book{Id: 1}).Data(book).Update() // do就不用指定具体字段
	if err != nil {
		req.Response.Write(result)
	} else {
		req.Response.Write(err)
	}
}

func (c *Hello) JoinTest(req *ghttp.Request) {
	empMd := dao.Emp.Ctx(req.Context())
	var emps []entity.Emp
	//err := empMd.Scan(&emps)
	err := empMd.With(entity.Dept{}).Scan(&emps) // with 表示连接dept
	if err == nil {
		req.Response.Writeln(emps)
	} else {
		req.Response.Writeln(err)
	}
}

func (c *Hello) OneToMany(req *ghttp.Request) {
	type myDept struct {
		g.Meta `orm:"table:dept"`
		Id     uint   `json:"id"`
		Name   string `json:"name"`
	}
	empMd := dao.Dept.Ctx(req.Context())
	var depts []entity.Dept
	err := empMd.With(entity.Emp{}).Scan(&depts) // with 表示连接dept

	if err == nil {
		req.Response.WriteJson(depts)
	} else {
		req.Response.Writeln(err)
	}
}

func (c *Hello) GetBookList(req *ghttp.Request) {
	list, err := service.Book().GetList(req.Context())
	if err == nil {
		req.Response.WriteJson(list)
	} else {
		req.Response.Writeln(err)
	}
	return
}

func (c *Hello) UploadFile(req *ghttp.Request) {
	file := req.GetUploadFile("file")
	if file != nil {
		req.Response.Writeln(file)
	}
}
