package service

import (
	"context"

	"github.com/website-pribadi/pkg/bookmark/domain/entity"
	"github.com/website-pribadi/pkg/bookmark/domain/repository"
)

type ReferenceServiceImpl struct {
	repo repository.TopicRepository
}

func NewReferenceService(repo repository.TopicRepository) ReferenceService {
	return &ReferenceServiceImpl{
		repo: repo,
	}
}

func (r ReferenceServiceImpl) FindMatchTopic(ctx context.Context, instance entity.Reference) (entity.Reference, error) {

	topic, err := r.repo.FindById(ctx, instance.TopicID)
	if err != nil {
		return entity.Reference{}, err
	}

	reference := entity.Reference{
		ID:      instance.ID,
		TopicID: topic.ID,
		Link:    instance.Link,
	}

	return reference, err
}
