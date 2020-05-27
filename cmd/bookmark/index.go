package cmd

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/go-kit/kit/log"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/website-pribadi/pkg/bookmark/domain/service"
	"github.com/website-pribadi/pkg/bookmark/infrastructure"
	"github.com/website-pribadi/pkg/bookmark/interface/endpoint"
	"github.com/website-pribadi/pkg/bookmark/interface/transport"
	"github.com/website-pribadi/pkg/bookmark/usecase"
)

type BookmarkServer struct {
	Ctx       context.Context
	Db        *sql.DB
	Histogram *prometheus.HistogramVec
	Logger    log.Logger
}

func (b *BookmarkServer) Start() http.Handler {

	var srv usecase.Usecase
	{
		TopicInfrastructure := infrastructure.TopicNewRepo(b.Db, b.Logger)
		service := service.NewService(TopicInfrastructure)

		ReferenceInfrastructure := infrastructure.ReferenceNewRepo(b.Db, b.Logger)

		middlwareUsecase := usecase.NewLoggingInterceptor(b.Logger, b.Histogram)                                    // setting up to insert middleware, type data of func
		srv = middlwareUsecase(usecase.NewUsecase(TopicInfrastructure, ReferenceInfrastructure, service, b.Logger)) // insert real function, call middleware func first
	}

	endpoint := endpoint.NewEndpoint(srv)
	handler := transport.NewHTTPServer(b.Ctx, endpoint, b.Logger)

	return handler
}