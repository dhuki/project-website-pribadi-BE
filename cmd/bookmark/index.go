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
		// these all using value receiver
		TopicInfrastructure := infrastructure.TopicNewRepo(b.Db, b.Logger)
		TopicService := service.NewService(TopicInfrastructure)
		ReferenceInfrastructure := infrastructure.ReferenceNewRepo(b.Db, b.Logger)
		ReferenceService := service.NewReferenceService(TopicInfrastructure)

		// if we're using pointer receiver class that implement interface, it must return pointer of class.
		// so there is two kind of implenet interface
		// by value, and by reference
		// if we're using value receiver class that implement interface, it must return value of class.
		// below example using pointer receiver
		usecaseImpl := &usecase.UsecaseImpl{
			TopicRepo:        TopicInfrastructure,
			ReferenceRepo:    ReferenceInfrastructure,
			TopicService:     TopicService,
			ReferenceService: ReferenceService,
			Logger:           b.Logger,
		}
		// end implemented

		middlewareUsecase := usecase.NewLoggingInterceptor(b.Logger, b.Histogram) // setting up to insert middleware, type data of func
		srv = middlewareUsecase(usecaseImpl)                                      // insert real function, call middleware func first
	}

	endpoint := endpoint.NewEndpoint(srv)
	handler := transport.NewHTTPServer(b.Ctx, endpoint, b.Logger)

	return handler
}
