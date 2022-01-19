package greet

import (
	"context"

	"github.com/yuchanns/template/utils"
	"github.com/yuchanns/template/vars"
)

type GreetDom struct {
	c    *utils.IoC
	repo GreetRepo
}

// NewGreetDom 在启动时调用
// 所以内部的错误可以直接 panic
func NewGreetDom(c *utils.IoC) *GreetDom {
	var repo GreetRepo
	c.MustInvoke(func(repository GreetRepo) {
		repo = repository
	})
	return &GreetDom{c: c, repo: repo}
}

func (g *GreetDom) SayHello(ctx context.Context, id string) (*Greeter, error) {
	gm, err := g.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return gm, nil
}

func (g *GreetDom) Async(ctx context.Context, name string, age int) error {
	id, err := g.repo.Create(ctx, name, age)
	if err != nil {
		return err
	}
	// 创建用户后需要做一些耗时的工作可以使用异步进行
	// 例如初始化用户的组织关系等
	vars.AsyncIDChan <- id
	return nil
}
