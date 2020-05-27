package service

import (
	"context"

	"github.com/website-pribadi/pkg/bookmark/domain/entity"
)

type TopicService interface {
	CreateTopic(ctx context.Context, topic entity.Topic) error
}
