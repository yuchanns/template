package startup

import (
	"github.com/yuchanns/template/internal/domain/async"
	"github.com/yuchanns/template/internal/domain/greet"
	"github.com/yuchanns/template/internal/infra/client"
	"github.com/yuchanns/template/internal/infra/mysql"
	"github.com/yuchanns/template/internal/server"
	"github.com/yuchanns/template/utils"
)

// RegisterContainer 注册容器实例
// 启动时调用 允许 panic
func RegisterContainer() *utils.IoC {
	c := utils.NewIoC()
	// 注册 Greet 服务
	c.MustProvide(func() (greet.GreetRepo, error) {
		return mysql.NewGreetImpl(), nil
	})
	c.MustProvide(func() (*greet.GreetDom, error) {
		return greet.NewGreetDom(c), nil
	})
	c.MustProvide(func() (*server.GreeterSrv, error) {
		return server.NewGreeterSrv(c), nil
	})

	// 注册事件服务
	c.MustProvide(func() (async.RelationRepo, error) {
		return client.NewUnknownImpl(c), nil
	})
	c.MustProvide(func() (*async.AsyncDom, error) {
		return async.NewAsyncDom(c), nil
	})
	c.MustProvide(func() (*server.AsyncSrv, error) {
		return server.NewAsyncSrv(c), nil
	})

	return c
}
