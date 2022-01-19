package client

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/yuchanns/template/internal/domain/async"
	"github.com/yuchanns/template/utils"
	"github.com/yuchanns/template/vars"
)

type UnknownImpl struct {
	c      *utils.IoC
	client *resty.Client
}

func NewUnknownImpl(c *utils.IoC) async.RelationRepo {
	client := resty.New()
	return &UnknownImpl{c: c, client: client}
}

func (u *UnknownImpl) InitUserRelation(ctx context.Context, id string) error {
	_, err := u.client.R().SetContext(ctx).
		SetPathParam("id", id).Get(vars.ServerUnknown + "/api/nowhere")
	return err
}
