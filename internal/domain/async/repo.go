package async

import "context"

type RelationRepo interface {
	InitUserRelation(context.Context, string) error
}
