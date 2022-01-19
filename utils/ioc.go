package utils

import "go.uber.org/dig"

type IoC struct {
	c *dig.Container
}

func NewIoC() *IoC {
	c := dig.New()
	return &IoC{c: c}
}

// MustProvide 用于启动时注册容器实例
// 错误将会停止进程直接 panic
func (c *IoC) MustProvide(constructor interface{}, opts ...dig.ProvideOption) {
	if err := c.c.Provide(constructor, opts...); err != nil {
		panic(err)
	}
}

// MustInvoke 用于启动时获取容器实例
// 错误将会停止进程直接 panic
func (c *IoC) MustInvoke(fn interface{}, opts ...dig.InvokeOption) {
	if err := c.c.Invoke(fn, opts...); err != nil {
		panic(err)
	}
}

// ProvideE 用于运行时注册容器实例
// 需要处理返回的错误
func (c *IoC) ProvideE(constructor interface{}, opts ...dig.ProvideOption) error {
	return c.c.Provide(constructor, opts...)
}

// InvokeE 用于运行时获取容器实例
// 需要处理返回的错误
func (c *IoC) InvokeE(fn interface{}, opts ...dig.InvokeOption) error {
	return c.c.Invoke(fn, opts...)
}
