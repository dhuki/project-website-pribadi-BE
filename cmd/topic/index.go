package cmd

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/go-kit/kit/log"

	"github.com/website-pribadi/pkg/topic/domain/service"
	"github.com/website-pribadi/pkg/topic/infrastructure"
	"github.com/website-pribadi/pkg/topic/interface/router"
	"github.com/website-pribadi/pkg/topic/usecase"
)

func NewService(ctx context.Context, db *sql.DB, logger log.Logger, histogram *prometheus.HistogramVec) http.Handler {
	var srv usecase.Usecase
	{
		infrastructure := infrastructure.NewRepo(db, logger)
		service := service.NewService(infrastructure)

		middlwareUsecase := usecase.NewLoggingInterceptor(logger, histogram)        // setting up to insert middleware first
		srv = middlwareUsecase(usecase.NewUseCase(infrastructure, service, logger)) // insert real function
	}

	endpoints := router.NewEndpoint(srv)
	handler := router.NewHTTPServer(ctx, endpoints)

	return handler
}
