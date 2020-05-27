package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/website-pribadi/pkg/bookmark/interface/model"
	"github.com/website-pribadi/pkg/bookmark/usecase"
)

type Endpoint struct {
	CreateTopic endpoint.Endpoint
	GetById     endpoint.Endpoint
	GetAllTopic endpoint.Endpoint
}

func NewEndpoint(usecase usecase.Usecase) Endpoint {
	return Endpoint{
		CreateTopic: makeCreateTopicEndpoint(usecase),
		GetById:     makeGetTopicByIdEndpoint(usecase),
		GetAllTopic: makeGetAllTopicEndpoint(usecase),
	}
}

func makeCreateTopicEndpoint(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.TopicRequest)
		response, err := usecase.CreateTopic(ctx, req)
		return response, err
	}
}

func makeGetTopicByIdEndpoint(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.TopicRequest)
		response, err := usecase.GetById(ctx, req)
		return response, err
	}
}

func makeGetAllTopicEndpoint(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		response, err := usecase.ListTopic(ctx)
		return response, err
	}
}
