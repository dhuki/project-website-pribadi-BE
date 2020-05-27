package usecase

import (
	"context"

	"github.com/go-kit/kit/log"
	guuid "github.com/google/uuid"
	"github.com/website-pribadi/pkg/bookmark/domain/entity"
	"github.com/website-pribadi/pkg/bookmark/domain/repository"
	"github.com/website-pribadi/pkg/bookmark/domain/service"
	"github.com/website-pribadi/pkg/bookmark/interface/model"
)

type UsecaseImpl struct {
	TopicRepo     repository.TopicRepository
	ReferenceRepo repository.ReferenceRepository
	service       service.Service
	logger        log.Logger
}

func NewUsecase(
	TopicRepo repository.TopicRepository,
	ReferenceRepo repository.ReferenceRepository,
	service service.Service, logger log.Logger) Usecase {
	return &UsecaseImpl{
		TopicRepo:     TopicRepo,
		ReferenceRepo: ReferenceRepo,
		service:       service,
		logger:        logger,
	}
}

func (t UsecaseImpl) CreateTopic(ctx context.Context, req model.TopicRequest) (model.BaseResponse, error) {
	// logger := log.With(t.logger, "method", "Create Topic")

	var response model.BaseResponse
	{
		instance := entity.Topic{
			ID:          guuid.New().String(),
			Name:        req.Name,
			Description: req.Description,
		}

		err := t.TopicRepo.CreateTopic(ctx, instance)
		if err != nil {
			response.Message = "Error"
			return response, err
		}

		response.Message = "Success"
	}

	// if isDuplicated, err := t.service.DuplicatedName(ctx, instance.Name); err != nil {
	// 	level.Error(logger).Log("err", err)
	// 	fmt.Println(isDuplicated)
	// }

	return response, nil
}

func (t UsecaseImpl) ListTopic(ctx context.Context) (model.BaseResponse, error) {

	var response model.BaseResponse
	{
		topics, err := t.TopicRepo.GetAllTopic(ctx)
		if err != nil {
			response.Message = "Error"
			return response, err
		}

		response.Data = topics
		response.Message = "Success"
	}

	return response, nil
}

func (t UsecaseImpl) GetById(ctx context.Context, req model.TopicRequest) (model.BaseResponse, error) {

	var response model.BaseResponse
	{
		topic, err := t.TopicRepo.FindById(ctx, req.ID)
		if err != nil {
			response.Message = "Error"
			return response, err
		}

		response.Data = topic
		response.Message = "Success"
	}

	return response, nil
}

func (t UsecaseImpl) CreateReference(ctx context.Context, req model.ReferenceRequest) (model.BaseResponse, error) {
	return model.BaseResponse{}, nil
}
