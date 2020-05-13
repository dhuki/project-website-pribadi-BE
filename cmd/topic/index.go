package cmd

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/go-kit/kit/log"

	"github.com/website-pribadi/pkg/topic/domain/service"
	"github.com/website-pribadi/pkg/topic/infrastructure"
	"github.com/website-pribadi/pkg/topic/interface/router"
	"github.com/website-pribadi/pkg/topic/usecase"
)

func NewService(ctx context.Context, db *sql.DB, logger log.Logger) http.Handler {
	var srv usecase.Usecase
	{
		infrastructure := infrastructure.NewRepo(db, logger)
		service := service.NewServices(infrastructure)

		srv = usecase.NewUseCase(infrastructure, service, logger)
	}

	endpoints := router.NewEndpoint(srv)
	handler := router.NewHTTPServer(ctx, endpoints)

	return handler
}
