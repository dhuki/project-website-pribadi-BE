package infrastructure

import (
	"context"
	"database/sql"

	"github.com/website-pribadi/pkg/bookmark/domain/repository"

	"github.com/website-pribadi/pkg/bookmark/domain/entity"

	"github.com/go-kit/kit/log"
)

type ReferenceRepoImpl struct {
	db     *sql.DB
	logger log.Logger
}

func ReferenceNewRepo(db *sql.DB, logger log.Logger) repository.ReferenceRepository {
	return &ReferenceRepoImpl{
		db:     db,
		logger: logger,
	}

}

func (r ReferenceRepoImpl) CreateReference(ctx context.Context, reference entity.Reference) error {
	return nil
}
