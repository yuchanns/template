package startup

import (
	"github.com/yuchanns/template/internal/server"
	"github.com/yuchanns/template/utils"
)

// RegisterEvent 用于注册事件
func RegisterEvent(c *utils.IoC) {
	{
		var eventSrv *server.AsyncSrv
		c.MustInvoke(func(srv *server.AsyncSrv) {
			eventSrv = srv
		})
		// 安全 goroutine
		go utils.SafeAsync(eventSrv.InitUserRelation)
	}
}
