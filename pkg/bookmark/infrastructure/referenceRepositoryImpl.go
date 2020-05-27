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
	sqlStatement := `
		INSERT INTO reference (id, topic_id, link)
		VALUES ($1, $2, $3)
	`

	// use db.Exec for operations that do not return rows (insert, delete, update)
	_, err := r.db.Exec(sqlStatement, reference.ID, reference.TopicID, reference.Link)
	if err != nil {
		return err
	}

	return nil
}
