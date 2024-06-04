package book

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
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
	cache := gcache.New()
	cache.Set(ctx, "k1", "v1", -1)
	g.Redis("cache1").Do(ctx, "SET", "k1", "v1")
	g2, err := g.Redis().Get(ctx, "k")
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	println(g2.String())
	return
}

func (s sBook) Update(ctx context.Context, book do.Book) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s sBook) Delete(ctx context.Context) (err error) {
	//TODO implement me
	panic("implement me")
}
