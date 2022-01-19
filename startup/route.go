package startup

import (
	"github.com/yuchanns/template/internal/server"
	"github.com/yuchanns/template/utils"

	"github.com/gin-gonic/gin"
)

// RegisterRoute 用于注册路由
func RegisterRoute(engine *gin.Engine, c *utils.IoC) {
	{
		var srv *server.GreeterSrv
		c.MustInvoke(func(httpSrv *server.GreeterSrv) {
			srv = httpSrv
		})
		engine.GET("/greet", utils.BuildGinHandler(srv.SayHello))
		engine.POST("/async", utils.BuildGinHandler(srv.Async))
	}
}
