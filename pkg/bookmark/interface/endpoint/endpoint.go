package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/website-pribadi/pkg/bookmark/interface/model"
	"github.com/website-pribadi/pkg/bookmark/usecase"
)

type Endpoint struct {
	CreateReferenceWithTopic endpoint.Endpoint
	CreateReference          endpoint.Endpoint
}

func NewEndpoint(usecase usecase.Usecase) Endpoint {
	return Endpoint{
		CreateReferenceWithTopic: makeCreateReferenceWithTopic(usecase),
		CreateReference:          makeCreateReferenceEndpoint(usecase),
	}
}

func makeCreateReferenceEndpoint(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.ReferenceRequest)
		response, err := usecase.CreateReference(ctx, req)
		return response, err
	}
}

func makeCreateReferenceWithTopic(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.ReferenceTopicRequest)
		response, err := usecase.CreateReferenceWithTopic(ctx, req)
		return response, err
	}
}
