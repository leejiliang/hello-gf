package book

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"hello-gf/internal/model/do"
	"hello-gf/internal/service"
)

type Book struct {
	msg string
}

func NewBook() *Book {
	return &Book{msg: "hello msg"}
}

func (c *Book) AddBook(req *ghttp.Request) {
	//list, err := service.Book().GetList(req.Context())
	err := service.Book().Add(req.Context(), do.Book{})
	if err != nil {
		return
	}
}
