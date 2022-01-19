package greet

import "context"

type GreetRepo interface {
	GetByID(context.Context, string) (*Greeter, error)
	Create(context.Context, string, int) (string, error)
}

type Greeter struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
