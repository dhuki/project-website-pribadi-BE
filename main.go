package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level" // side-effect only run func init inside its library
	_ "github.com/lib/pq"             // side-effect only run func init inside its library

	topicService "github.com/website-pribadi/cmd/topic"
	"github.com/website-pribadi/config"
)

// db heroku
// const dbURI = "postgres://ijjfgiyjcbrznd:028ccc705a477aabf05483a69471b3cd349122d598495748014d446c7aad41dd@ec2-174-129-255-21.compute-1.amazonaws.com:5432/d4ppbeeimehina"

func main() {

	var httpAddr = flag.String("http", ":8080", "http listen address")
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"project", "website pribadi",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	// 0666 is chmod meaning permission in that file (0666 -> can write (write something to the file) /read (see the file) )
	// os.O_CREATE if it's not found | os.O_APPEND if it's found
	f, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		level.Error(logger).Log("exit", err)
		f.WriteString(level.Error(logger).Log("exit", err).Error())
		os.Exit(-1) // exit program with status -1
	}

	var db *sql.DB
	{
		database := config.NewDatabase(logger, f)
		db = database.Start()
	}

	ctx := context.Background()

	errs := make(chan error)

	// go routine
	go func() {
		c := make(chan os.Signal, 1) //make(T, n) T is type and n is length
		// registers the given channel to receive notifications of the specified signals.
		// SIGINT (Signal Interrupt (CTRL + C))
		// SIGTERM (Signal Terminated (KILL command))
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c) // it's blocking until receiver a signal
	}()

	// go routine
	go func() {
		fmt.Println("listening on port", *httpAddr)

		mux := http.NewServeMux()
		mux.Handle("/topic/", http.StripPrefix("/topic/api", topicService.NewService(ctx, db, logger)))

		errs <- http.ListenAndServe(*httpAddr, mux) // it's blocking until error emerge while listen to webserver
	}()

	level.Error(logger).Log("exit", <-errs)

}
