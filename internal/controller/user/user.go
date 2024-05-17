package user

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"hello-gf/api/user"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) AddUser(req *ghttp.Request) {
	req.Response.Writeln("添加用户")
}

func (c *Controller) DelUser(req *ghttp.Request) {
	req.Response.Writeln("删除用户")
}

func (c *Controller) ModifyUser(req *ghttp.Request) {
	req.Response.Writeln("修改用户")
}

func (c *Controller) QueryUser(req *ghttp.Request) {
	req.Response.Writeln("查询用户")
}

func (c *Controller) Add2(ctx context.Context, req *user.AddReq) (resp *user.AddRes, err error) {
	resp = &user.AddRes{
		Name: "张三",
		Age:  20,
	}

	//g.RequestFromCtx(ctx).Response.Writeln("Hello add2")
	return
}
