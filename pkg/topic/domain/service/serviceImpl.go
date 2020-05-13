package service

import (
	"context"

	"github.com/website-pribadi/pkg/topic/domain/repository"
)

type TopicServicesImpl struct {
	repo repository.Repository
}

func NewServices(repo repository.Repository) Service {
	return &TopicServicesImpl{
		repo: repo,
	}
}

func (t *TopicServicesImpl) DuplicatedName(ctx context.Context, name string) (bool, error) {

	isDuplicated, err := t.repo.FindByName(ctx, name)
	if err != nil {
		return isDuplicated, err
	}

	return isDuplicated, nil
}
