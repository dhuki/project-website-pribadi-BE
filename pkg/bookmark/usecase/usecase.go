package usecase

import (
	"context"

	"github.com/website-pribadi/pkg/bookmark/interface/model"
)

type Usecase interface {
	CreateReferenceWithTopic(ctx context.Context, req model.ReferenceTopicRequest) (model.BaseResponse, error)
	CreateReference(ctx context.Context, req model.ReferenceRequest) (model.BaseResponse, error)
}
