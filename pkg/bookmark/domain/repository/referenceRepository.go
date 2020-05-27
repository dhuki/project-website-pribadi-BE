package repository

import (
	"context"

	"github.com/website-pribadi/pkg/bookmark/domain/entity"
)

type ReferenceRepository interface {
	CreateReference(ctx context.Context, reference entity.Reference) error
}
