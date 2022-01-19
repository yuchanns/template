package mysql

import (
	"context"

	"github.com/yuchanns/template/internal/domain/greet"
	"github.com/yuchanns/template/internal/infra/mysql/models"
	"github.com/yuchanns/template/utils"
	"github.com/yuchanns/template/vars"
)

type GreetImpl struct{}

func NewGreetImpl() greet.GreetRepo {
	return &GreetImpl{}
}

func (g *GreetImpl) GetByID(ctx context.Context, id string) (*greet.Greeter, error) {
	gm := &models.Greet{}
	if err := vars.DB.WithContext(ctx).Where("id = ?", id).Find(&gm).Error; err != nil {
		return nil, err
	}

	return ModelToGreeter(gm)
}

func (g *GreetImpl) Create(ctx context.Context, name string, age int) (string, error) {
	gm := &models.Greet{
		ID:   vars.SnowflakeNode.Generate().String(),
		Name: name,
		Age:  age,
	}
	if err := vars.DB.WithContext(ctx).Create(gm).Error; err != nil {
		return "", err
	}

	return gm.ID, nil
}

func ModelToGreeter(in *models.Greet) (*greet.Greeter, error) {
	out := &greet.Greeter{}
	err := utils.Convert(in, &out)
	return out, err
}
