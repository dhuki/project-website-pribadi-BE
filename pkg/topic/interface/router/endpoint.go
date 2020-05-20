package router

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/website-pribadi/pkg/topic/interface/transport"
	"github.com/website-pribadi/pkg/topic/usecase"
)

type Endpoints struct {
	CreateTopic endpoint.Endpoint
	GetById     endpoint.Endpoint
}

func NewEndpoint(usecase usecase.Usecase) Endpoints {
	return Endpoints{
		CreateTopic: makeCreateTopicEndpoint(usecase),
		//ListTopic : makeCreateTopicEndpoint(t)
	}
}

func makeCreateTopicEndpoint(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.Request)
		topic, err := usecase.CreateTopic(ctx, req)
		return transport.Response{
			ID:          topic.ID,
			Name:        topic.Name,
			Description: topic.Description,
		}, err
	}
}

func makeGetTopicByIdEndpoint(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.Request)
		topic, err := usecase.GetById(ctx, req)
		return transport.Response{
			ID:          topic.ID,
			Name:        topic.Name,
			Description: topic.Description,
		}, err
	}
}
