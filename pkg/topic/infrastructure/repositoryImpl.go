package infrastructure

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/website-pribadi/pkg/topic/domain/entity"
	"github.com/website-pribadi/pkg/topic/domain/repository"
)

type TopicRepoImpl struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) repository.Repository {
	return &TopicRepoImpl{
		db:     db,
		logger: log.With(logger, "Repo lecturer", "sql"),
	}
}

func (t TopicRepoImpl) GetAllTopic(ctx context.Context) ([]entity.Topic, error) {
	rows, err := t.db.Query("SELECT * FROM TOPIC")
	defer rows.Close() //will execute after all line in scope function executed
	if err != nil {
		return nil, err
	}

	var topics []entity.Topic

	// check if query has next value
	for rows.Next() {

		instance := &entity.Topic{}

		if topics == nil {
			topics = []entity.Topic{} //define a slice -> FYI slice more lightweight than array because its using pointer
		}

		err := rows.Scan(&instance.ID, &instance.Name, &instance.Description) //Scan copies the columns from the matched row (in database) into the values pointed (in struct) at by its destination
		if err != nil {
			return nil, err
		}

		topics = append(topics, *instance)
	}

	return topics, nil
}

func (t TopicRepoImpl) CreateTopic(ctx context.Context, topic entity.Topic) error {
	rows, err := t.db.Exec(`
			SELECT * FROM TOPIC 
			WHERE LOWER(TRIM(name)) == $1
		`, topic.Name)

	numRows, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println(numRows)

	return nil
}

func (t TopicRepoImpl) FindByName(ctx context.Context, name string) (bool, error) {

	isDuplicated := false

	rows, err := t.db.Exec(`
			SELECT * FROM TOPIC 
			WHERE LOWER(TRIM(name)) == $1
		`, name)
	if err != nil {
		return false, err
	}

	numRows, err := rows.RowsAffected()
	if err != nil {
		return false, err
	}

	if numRows > 0 {
		isDuplicated = true
	}

	return isDuplicated, nil
}

func (t TopicRepoImpl) FindById(ctx context.Context, id string) (entity.Topic, error) {
	instance := &entity.Topic{}

	// in postgres we are using $1 for parameter
	err := t.db.QueryRow(
		`SELECT * FROM TOPIC WHERE id = $1`, id).Scan(&instance.ID, &instance.Name, &instance.Description)
	if err != nil {
		return entity.Topic{}, err
	}
	return *instance, nil
}
