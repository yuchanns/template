package async

import (
	"context"
	"fmt"

	"github.com/yuchanns/template/utils"
)

type AsyncDom struct {
	c    *utils.IoC
	repo RelationRepo
}

func NewAsyncDom(c *utils.IoC) *AsyncDom {
	var repo RelationRepo
	c.MustInvoke(func(repository RelationRepo) {
		repo = repository
	})
	return &AsyncDom{c: c, repo: repo}
}

func (a *AsyncDom) InitUserRelation(ctx context.Context, id string) {
	// 耗时任务，初始化用户关系等工作
	if err := a.repo.InitUserRelation(ctx, id); err != nil {
		fmt.Println(err)
	}
}
