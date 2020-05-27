package service

import (
	"context"

	"github.com/website-pribadi/pkg/bookmark/domain/entity"
)

type ReferenceService interface {
	FindMatchTopic(context.Context, entity.Reference) (entity.Reference, error)
}
