package server

import (
	"github.com/yuchanns/template/internal/domain/greet"
	"github.com/yuchanns/template/utils"

	"github.com/gin-gonic/gin"
)

// GreeterSrv 是 Greeter 服务入口
type GreeterSrv struct {
	c   *utils.IoC
	dom *greet.GreetDom
}

// NewGreeterSrv 是 Greeter 构造函数
// 在启动时调用 允许错误 panic
func NewGreeterSrv(c *utils.IoC) *GreeterSrv {
	var dom *greet.GreetDom
	c.MustInvoke(func(domain *greet.GreetDom) {
		dom = domain
	})
	return &GreeterSrv{c: c, dom: dom}
}

func (g *GreeterSrv) SayHello(ctx *gin.Context) (interface{}, error) {
	req := RequestHello{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		return nil, err
	}
	gm, err := g.dom.SayHello(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	resp := &ResponseHello{}
	err = utils.Convert(gm, &resp)
	return resp, err
}

func (g *GreeterSrv) Async(ctx *gin.Context) (interface{}, error) {
	req := RequestAsync{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	g.dom.Async(ctx, req.Name, req.Age)
	return nil, nil
}

type RequestHello struct {
	ID string `form:"id"`
}

type ResponseHello struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type RequestAsync struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
