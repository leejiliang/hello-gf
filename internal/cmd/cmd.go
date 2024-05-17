package cmd

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"hello-gf/internal/controller/hello"
	"hello-gf/internal/controller/user"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

func gf2(req *ghttp.Request) {
	req.Response.Write("hello GF2")
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// 路由注册
			s.BindHandler("/my-hello", func(req *ghttp.Request) {
				req.Response.Write("hello GF")
			})

			s.BindHandler("/my-hello2", gf2)

			hello := hello.NewHello()
			s.BindHandler("GET:/say", hello.SayHello)

			// 批量绑定所有controller中的方法
			//s.BindObject("/user", user.New())
			// 绑定指定的方法, 多个方法名之间用逗号分隔。
			//s.BindObject("/user", user.New(), "AddUser,DelUser")
			// 绑定某一个方法
			//s.BindObjectMethod("/user/query", user.New(), "QueryUser")
			// 根据Rest规范来选择绑定方法, 例如： GET POST PUT DELETE 方法名需要固定
			s.BindObjectRest("/user", user.New())
			//s.Group("/", func(group *ghttp.RouterGroup) {
			//	group.Middleware(ghttp.MiddlewareHandlerResponse)
			//	group.Bind(
			//		hello.NewV1(),
			//	)
			//})
			s.Run()
			return nil
		},
	}
)
