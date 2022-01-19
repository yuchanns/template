package server

import (
	"context"
	"time"

	"github.com/panjf2000/ants/v2"
	"github.com/yuchanns/template/internal/domain/async"
	"github.com/yuchanns/template/utils"
	"github.com/yuchanns/template/vars"
)

type AsyncSrv struct {
	c   *utils.IoC
	dom *async.AsyncDom
}

func NewAsyncSrv(c *utils.IoC) *AsyncSrv {
	var dom *async.AsyncDom
	c.MustInvoke(func(domain *async.AsyncDom) {
		dom = domain
	})
	return &AsyncSrv{c: c, dom: dom}
}

// InitUserRelation 初始化用户关系
// 这是个常驻的 goroutine
func (a *AsyncSrv) InitUserRelation() {
	parentCtx := context.Background()
	p, _ := ants.NewPool(10)
	defer p.Release()
	for {
		id := <-vars.AsyncIDChan
		// 1. 顺序消费
		// 如果异步任务太多会造成 channel 阻塞等待
		// 这种情况考虑将 channel 改造成可动态扩容
		// 参考 https://colobu.com/2021/05/11/unbounded-channel-in-go/
		/* ctx, cancel := context.WithTimeout(parentCtx, time.Second)
		a.dom.InitUserRelation(ctx, id)
		cancel() */

		// 2. 异步消费
		// 使用具有 recover 能力的 goroutine
		// 在 panic 时可以记录错误日志
		// 拿到就消费，可能存在 goroutine 暴涨的情况
		/* go utils.SafeAsync(func() {
			// 超时控制手段
			ctx, cancel := context.WithTimeout(parentCtx, time.Second)
			defer cancel()
			a.dom.InitUserRelation(ctx, id)
		}) */

		// 3. 限量异步消费
		// 使用具有 recover 能力的 goroutine
		// 同时限制最大并发量，避免 goroutine 过多
		p.Submit(utils.SafeAsyncFunc(func() {
			// 超时控制手段
			ctx, cancel := context.WithTimeout(parentCtx, time.Second)
			defer cancel()
			a.dom.InitUserRelation(ctx, id)
		}))
	}
}
