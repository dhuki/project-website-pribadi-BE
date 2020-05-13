package infrastructure

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/website-pribadi/pkg/topic/domain/repository"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/website-pribadi/config"
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

	// 0666 is chmod meaning permission in that file (0666 -> can write (write something to the file) /read (see the file) )
	// os.O_CREATE if it's not found | os.O_APPEND if it's found
	f, err := os.OpenFile("logfile_testing.log", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		level.Error(logger).Log("exit", err)
		f.WriteString(level.Error(logger).Log("exit", err).Error())
		os.Exit(-1) // exit program with status -1
	}

	var db *sql.DB
	{
		database := config.NewDatabase(logger, f)
		db = database.Start(os.ExpandEnv("$GOPATH\\src\\github.com\\website-pribadi\\.env"))
	}

	repo = NewRepo(db, logger)

	code := m.Run()

	os.Exit(code)
}

func TestFindById(t *testing.T) {
	topic, err := repo.FindById(context.TODO(), "asu")
	if err != nil {
		t.Errorf("error occur %d", err)
	}

	if topic.ID == "asu" {
		t.Errorf("value is null %v", topic.ID)
	}
}
