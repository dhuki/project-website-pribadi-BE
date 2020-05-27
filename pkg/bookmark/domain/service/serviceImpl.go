package service

import (
	"github.com/website-pribadi/pkg/bookmark/domain/repository"
)

type TopicServicesImpl struct {
	repo repository.TopicRepository
}

func NewService(repo repository.TopicRepository) Service {
	return &TopicServicesImpl{
		repo: repo,
	}
}

// func (t *TopicServicesImpl) CreateTopic(ctx context.Context, name string) (bool, error) {

// 	isDuplicated, err := t.repo.FindByName(ctx, name)
// 	if err != nil {
// 		return isDuplicated, err
// 	}

// 	return isDuplicated, nil
// }
