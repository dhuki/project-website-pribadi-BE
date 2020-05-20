package infrastructure

// testing with BDD
// if we want to see coverage testing make file coverage : ginkgo -cover -coverprofile=coverage.out
// see coverage with html : go tool cover -html=coverage.out

import (
	"database/sql"
	"fmt"
	"go/build"
	"os"
	"testing"

	"github.com/go-kit/kit/log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/website-pribadi/config"
	"github.com/website-pribadi/pkg/topic/domain/repository"
)

var repo repository.Repository

func TestInfrastructure(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Infrastructure Suite")
}

// set up database once before the entire test suite
var _ = BeforeSuite(func() {
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

	var db *sql.DB
	{
		source := fmt.Sprintf("%v\\src\\github.com\\website-pribadi\\.env", build.Default.GOPATH)
		database := config.NewDatabase(logger)
		db = database.Start(source)
	}

	repo = NewRepo(db, logger)
})
