package main

import (
	"github.com/yuchanns/template/startup"

	"github.com/gin-gonic/gin"
)

func main() {
	// 实际上这里我们一般不会直接使用 http server (例如 gin)
	// 以及手动调用这些注册
	// 而是封装成一个 Application 结构体
	// 在 Run 的时候自动去完成
	// 1. http server 和 event server 的实例化，中间件注册
	// 2. 启动时注册
	// 3. 分别在独立的 goroutine 里启动 http server 和 event server
	// 4. 安装信号监听器，阻塞主 goroutine 并实现优雅关机
	// 这里为了方便理解省略这些细节
	engine := gin.Default()

	// 一系列的启动时注册
	c := startup.RegisterContainer()
	startup.RegisterVars()
	startup.RegisterRoute(engine, c)
	startup.RegisterEvent(c)

	// 启动服务
	engine.Run()
}
