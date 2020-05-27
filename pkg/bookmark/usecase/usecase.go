package usecase

import (
	"context"

	"github.com/website-pribadi/pkg/bookmark/interface/model"
)

type Usecase interface {
	CreateTopic(ctx context.Context, req model.TopicRequest) (model.BaseResponse, error)
	ListTopic(ctx context.Context) (model.BaseResponse, error)
	GetById(ctx context.Context, req model.TopicRequest) (model.BaseResponse, error)
	CreateReference(ctx context.Context, req model.ReferenceRequest) (model.BaseResponse, error)
}
