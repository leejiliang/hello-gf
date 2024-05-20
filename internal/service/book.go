package service

import (
	"context"
	"hello-gf/internal/model/do"
	"hello-gf/internal/model/entity"
)

// 1. 定义接口
type IBook interface {
	GetList(ctx context.Context) (books []entity.Book, err error)
	Add(ctx context.Context, book do.Book) (err error)
	Update(ctx context.Context, book do.Book) (err error)
	Delete(ctx context.Context) (err error)
}

// 2. 定义接口变量
var localBook IBook

// 3. 定义获取接口实例的函数
func Book() IBook {
	if localBook == nil {
		panic("IBook 接口未实现或未注册")
	}
	return localBook
}

// 4. 定义接口实现的注册方法
func RegisterBook(i IBook) {
	localBook = i
}
