package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"go/build"
	"os"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/website-pribadi/config"
	"github.com/website-pribadi/pkg/topic/domain/repository"
)

var repo repository.Repository

func TestMain(m *testing.M) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"project", "website pribadi testing",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started testing")
	defer level.Info(logger).Log("msg", "service ended testing")

	var db *sql.DB
	{
		source := fmt.Sprintf("%v\\src\\github.com\\website-pribadi\\.env", build.Default.GOPATH)
		database := config.NewDatabase(logger)
		db = database.Start(source)
	}

	repo = NewRepo(db, logger)

	fmt.Println("Do stuff BEFORE the tests!")

	code := m.Run()

	fmt.Println("Do stuff AFTER the tests!")

	os.Exit(code)
}

func TestFindById(t *testing.T) {
	topic, err := repo.FindById(context.TODO(), "asu")
	if err != nil {
		t.Errorf("error occur %d", err)
	}

	if topic.ID != "asu" {
		t.Errorf("value is null %v", topic.ID)
	}
}
