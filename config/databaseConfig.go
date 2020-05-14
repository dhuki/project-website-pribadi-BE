package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // to get driver from psql
)

type postgresDB interface {
	Start(env string) *sql.DB
}

type postgresDBImpl struct {
	logger log.Logger
}

func NewDatabase(logger log.Logger) postgresDB {
	return &postgresDBImpl{
		logger: logger,
	}
}

func (p postgresDBImpl) Start(env string) *sql.DB {
	var db *sql.DB
	{
		err := godotenv.Load(env)
		if err != nil {
			level.Error(p.logger).Log("exit", err)
		}

		username := os.Getenv("db.username")
		password := os.Getenv("db.password")
		dbName := os.Getenv("db.name")
		dbHost := os.Getenv("db.host")
		// port is default since we use library from psql

		dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
		fmt.Println(dbURI)

		db, err = sql.Open("postgres", dbURI)
		if err != nil {
			level.Error(p.logger).Log("exit", err)
			os.Exit(-1)
		}
	}
	return db
}
