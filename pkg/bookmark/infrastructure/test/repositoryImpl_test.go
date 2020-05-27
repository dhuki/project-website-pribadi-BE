package infrastructure

// testing using TDD

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"go/build"
// 	"os"
// 	"reflect"
// 	"testing"

// 	"github.com/go-kit/kit/log"
// 	"github.com/go-kit/kit/log/level"
// 	guuid "github.com/google/uuid"
// 	"github.com/website-pribadi/config"
// 	"github.com/website-pribadi/pkg/topic/domain/entity"
// 	"github.com/website-pribadi/pkg/topic/domain/repository"
// )

// var repo repository.Repository

// // configuration before run another func test
// func TestMain(m *testing.M) {
// 	var logger log.Logger
// 	{
// 		logger = log.NewLogfmtLogger(os.Stderr)
// 		logger = log.NewSyncLogger(logger)
// 		logger = log.With(logger,
// 			"project", "website pribadi testing",
// 			"time:", log.DefaultTimestampUTC,
// 			"caller", log.DefaultCaller,
// 		)
// 	}

// 	level.Info(logger).Log("msg", "service started testing")
// 	defer level.Info(logger).Log("msg", "service ended testing")

// 	var db *sql.DB
// 	{
// 		source := fmt.Sprintf("%v\\src\\github.com\\website-pribadi\\.env", build.Default.GOPATH)
// 		database := config.NewDatabase(logger)
// 		db = database.Start(source)
// 	}

// 	repo = NewRepo(db, logger)

// 	fmt.Println("Do stuff BEFORE the tests!")

// 	code := m.Run()

// 	fmt.Println("Do stuff AFTER the tests!")

// 	os.Exit(code)
// }

// func TestFindById(t *testing.T) {
// 	topic, err := repo.FindById(context.TODO(), "asu")
// 	if err != nil {
// 		t.Errorf("error occur %d", err) // %d base 10 -> int64
// 	}

// 	if topic.ID != "asu" {
// 		t.Errorf("value is null %v", topic.ID) // %v default value format all type can user this
// 	}
// }

// func TestGetAllTopic(t *testing.T) {
// 	expected := []entity.Topic{ // insert into slice of object topic
// 		entity.Topic{
// 			ID:          "asu",
// 			Name:        "asu",
// 			Description: "asu",
// 		},
// 	}

// 	value, err := repo.GetAllTopic(context.TODO())
// 	if err != nil {
// 		t.Errorf("error occur %d", err)
// 	}

// 	if len(value) != len(expected) {
// 		t.Errorf("error occur mismatched length of array expected : %b, actual : %b", len(expected), len(value)) // %b base2 int16
// 	}

// 	// compare two list of topic slice
// 	if !reflect.DeepEqual(value, expected) {
// 		t.Errorf("error occur value is not same as expected : %v, actual : %v", expected, value)
// 	}
// }

// func TestCreateTopic(t *testing.T) {
// 	expected := entity.Topic{
// 		ID:          guuid.New().String(),
// 		Name:        "",
// 		Description: "",
// 	}

// 	err := repo.CreateTopic(context.TODO(), expected)
// 	if err != nil {
// 		t.Errorf("error occur %d", err)
// 	}

// 	actual, err := repo.FindById(context.TODO(), expected.ID)
// 	if err != nil {
// 		t.Errorf("error occur %d", err)
// 	}

// 	if actual == expected {

// 	}

// }
