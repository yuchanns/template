package main

import (
	"github.com/yuchanns/template/startup"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// 一系列的启动时注册
	c := startup.RegisterContainer()
	startup.RegisterVars()
	startup.RegisterRoute(engine, c)
	startup.RegisterEvent(c)

	// 启动服务
	engine.Run()
}
