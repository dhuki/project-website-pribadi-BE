package service

import (
	"github.com/website-pribadi/pkg/topic/domain/repository"
)

type TopicServicesImpl struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
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
