package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/website-pribadi/pkg/bookmark/interface/endpoint"
	"github.com/website-pribadi/pkg/bookmark/interface/model"

	transportKit "github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoint, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),                            // custom response error to client
		httptransport.ServerErrorHandler(transportKit.NewLogErrorHandler(logger)), // log error to terminal
	}

	r.Methods("POST").Path("/").Handler(httptransport.NewServer(
		endpoints.CreateTopic,
		model.DecodeTopicReq,
		model.EncodeResponse,
	))

	r.Methods("GET").Path("/{id}").Handler(httptransport.NewServer(
		endpoints.GetById,
		model.DecodeTopicGetById,
		model.EncodeResponse,
		options...,
	))

	r.Methods("GET").Path("/").Handler(httptransport.NewServer(
		endpoints.GetAllTopic,
		model.DecodeTopicGetAll,
		model.EncodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(model.BaseResponse{
		Message: err.Error(),
	})
}
