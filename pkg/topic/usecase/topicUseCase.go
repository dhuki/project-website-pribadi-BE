package usecase

import (
	"context"

	"github.com/website-pribadi/pkg/topic/domain/entity"
	"github.com/website-pribadi/pkg/topic/interface/transport"
)

type Usecase interface {
	CreateTopic(ctx context.Context, req transport.Request) (entity.Topic, error)
	ListTopic(ctx context.Context) ([]entity.Topic, error)
	GetById(ctx context.Context, req transport.Request) (entity.Topic, error)
}
