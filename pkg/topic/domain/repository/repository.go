package repository

import (
	"context"

	"github.com/website-pribadi/pkg/topic/domain/entity"
)

type Repository interface {
	GetAllTopic(ctx context.Context) ([]entity.Topic, error)
	CreateTopic(ctx context.Context, topic entity.Topic) error
	FindByName(ctx context.Context, name string) (bool, error)
	FindById(ctx context.Context, id string) (entity.Topic, error)
}
