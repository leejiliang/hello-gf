package user

import "github.com/gogf/gf/v2/net/ghttp"

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
