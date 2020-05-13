package usecase

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/website-pribadi/pkg/topic/domain/entity"
	"github.com/website-pribadi/pkg/topic/domain/repository"
	"github.com/website-pribadi/pkg/topic/domain/service"
	"github.com/website-pribadi/pkg/topic/interface/transport"
)

type TopicUseCaseImpl struct {
	repo    repository.Repository
	service service.Service
	logger  log.Logger
}

func NewUseCase(repo repository.Repository, service service.Service, logger log.Logger) Usecase {
	return &TopicUseCaseImpl{
		repo:    repo,
		service: service,
		logger:  logger,
	}
}

func (t TopicUseCaseImpl) CreateTopic(ctx context.Context, req transport.Request) (entity.Topic, error) {
	logger := log.With(t.logger, "method", "Create Topic")

	instance := entity.Topic{
		Name:        req.Name,
		Description: req.Description,
	}

	if isDuplicated, err := t.service.DuplicatedName(ctx, instance.Name); err != nil {
		level.Error(logger).Log("err", err)
		fmt.Println(isDuplicated)
	}

	return entity.Topic{}, nil

}

func (t TopicUseCaseImpl) ListTopic(ctx context.Context) ([]entity.Topic, error) {
	return []entity.Topic{}, nil
}

func (t TopicUseCaseImpl) GetById(ctx context.Context, req transport.Request) (entity.Topic, error) {
	// logger := log.With(t.logger, "method", "Get By Id")

	return entity.Topic{}, nil
}
