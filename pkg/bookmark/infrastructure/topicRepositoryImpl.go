package infrastructure

import (
	"context"
	"database/sql"

	"github.com/go-kit/kit/log"
	"github.com/website-pribadi/pkg/bookmark/domain/entity"
	"github.com/website-pribadi/pkg/bookmark/domain/repository"
)

type TopicRepoImpl struct {
	db     *sql.DB
	logger log.Logger
}

func TopicNewRepo(db *sql.DB, logger log.Logger) repository.TopicRepository {
	return &TopicRepoImpl{
		db:     db,
		logger: log.With(logger, "Repo lecturer", "sql"),
	}
}

func (t TopicRepoImpl) CreateTopic(ctx context.Context, topic entity.Topic) error {
	sqlStatement := `
		INSERT INTO topic (id, name, description)
		VALUES ($1, $2, $3)
	`

	// use db.Exec for operations that do not return rows (insert, delete, update)
	_, err := t.db.Exec(sqlStatement, topic.ID, topic.Name, topic.Description)
	if err != nil {
		return err
	}

	// _, err := result.RowsAffected() // check row affected by query result
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (t TopicRepoImpl) GetAllTopic(ctx context.Context) ([]entity.Topic, error) {
	// use db.Query for return a set of rows
	rows, err := t.db.Query("SELECT * FROM topic")
	defer rows.Close() // will execute after all line in scope function executed
	if err != nil {
		return nil, err
	}

	var topics []entity.Topic

	// check if query has next value
	for rows.Next() {
		instance := &entity.Topic{}

		if topics == nil {
			topics = []entity.Topic{} // define a slice -> FYI slice more lightweight than array because its using pointer
		}

		err := rows.Scan(&instance.ID, &instance.Name, &instance.Description) // Scan copies the columns from the matched row (in database) into the values pointed (in struct) at by its destination
		if err != nil {
			return nil, err
		}

		topics = append(topics, *instance)
	}

	return topics, nil
}

func (t TopicRepoImpl) FindById(ctx context.Context, id string) (entity.Topic, error) {
	instance := &entity.Topic{}

	// in postgres we are using $1 for parameter
	// db.QueryRow is used to get at most one return
	err := t.db.QueryRow(
		`SELECT * FROM topic WHERE id = $1`, id).Scan(&instance.ID, &instance.Name, &instance.Description)
	if err != nil {
		return entity.Topic{}, err
	}

	return *instance, nil
}
