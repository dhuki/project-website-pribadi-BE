package service

import (
	"context"

	"github.com/website-pribadi/pkg/bookmark/domain/entity"
	"github.com/website-pribadi/pkg/bookmark/domain/repository"
)

type TopicServicesImpl struct {
	repo repository.TopicRepository
}

func NewService(repo repository.TopicRepository) TopicService {
	return &TopicServicesImpl{
		repo: repo,
	}
}

func (t *TopicServicesImpl) CreateTopic(ctx context.Context, topic entity.Topic) error {
	return nil
}
