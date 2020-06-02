package usecase

import (
	"context"
	"database/sql"

	"github.com/go-kit/kit/log"
	guuid "github.com/google/uuid"
	"github.com/website-pribadi/common"
	"github.com/website-pribadi/pkg/bookmark/domain/entity"
	"github.com/website-pribadi/pkg/bookmark/domain/repository"
	"github.com/website-pribadi/pkg/bookmark/domain/service"
	"github.com/website-pribadi/pkg/bookmark/interface/model"
)

type UsecaseImpl struct {
	TopicRepo        repository.TopicRepository
	ReferenceRepo    repository.ReferenceRepository
	TopicService     service.TopicService
	ReferenceService service.ReferenceService
	Logger           log.Logger
}

func (t *UsecaseImpl) CreateReferenceWithTopic(ctx context.Context, req model.ReferenceTopicRequest) (model.BaseResponse, error) {
	// logger := log.With(t.logger, "method", "Create Topic")

	// go func() (model.BaseResponse, error) {
	// 	select {
	// 	case <-ctx.Done():
	// 		return model.BaseResponse{}, ctx.Err()
	// 	}
	// }()

	var response model.BaseResponse
	{
		instanceTopic := entity.Topic{
			ID:          guuid.New().String(),
			Name:        req.NameTopic,
			Description: req.DescriptionTopic,
		}

		err := t.TopicRepo.CreateTopic(ctx, instanceTopic)
		if err != nil {
			return model.BaseResponse{}, err
		}

		for _, value := range req.Links {
			instanceReference := entity.Reference{
				ID:      guuid.New().String(),
				TopicID: instanceTopic.ID,
				Link:    value,
			}

			t.ReferenceRepo.CreateReference(ctx, instanceReference)
			if err != nil {
				return model.BaseResponse{}, err
			}
		}

		response.Message = "Success"
	}

	// if isDuplicated, err := t.service.DuplicatedName(ctx, instance.Name); err != nil {
	// 	level.Error(logger).Log("err", err)
	// 	fmt.Println(isDuplicated)
	// }

	return response, nil
}

func (t *UsecaseImpl) CreateReference(ctx context.Context, req model.ReferenceRequest) (model.BaseResponse, error) {

	var response model.BaseResponse
	{
		instance := entity.Reference{
			ID:      guuid.New().String(),
			TopicID: req.TopicID,
			Link:    req.Link,
		}

		reference, err := t.ReferenceService.FindMatchTopic(ctx, instance)

		switch err {
		case sql.ErrNoRows:
			return model.BaseResponse{
				Message: common.ErrNotFound.Error(),
				ErrCode: common.ErrNotFoundCode,
			}, nil
		case err:
			return model.BaseResponse{}, nil
		}

		response.Data = reference
	}

	return response, nil
}
