package repository

import (
	"context"

	"github.com/website-pribadi/pkg/bookmark/domain/entity"
)

type TopicRepository interface {
	GetAllTopic(ctx context.Context) ([]entity.Topic, error)
	CreateTopic(ctx context.Context, topic entity.Topic) error
	FindById(ctx context.Context, id string) (entity.Topic, error)
}
