package book

import (
	"context"
	"hello-gf/internal/dao"
	"hello-gf/internal/model/do"
	"hello-gf/internal/model/entity"
	"hello-gf/internal/service"
)

func init() {
	service.RegisterBook(&sBook{})
}

type sBook struct {
}

func (s sBook) GetList(ctx context.Context) (books []entity.Book, err error) {
	err = dao.Book.Ctx(ctx).Scan(&books)
	//TODO implement me
	return
	//panic("implement me")
}

func (s sBook) Add(ctx context.Context, book do.Book) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s sBook) Update(ctx context.Context, book do.Book) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s sBook) Delete(ctx context.Context) (err error) {
	//TODO implement me
	panic("implement me")
}
