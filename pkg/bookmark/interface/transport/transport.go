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

	r.Methods("POST").Path("/topic").Handler(adminOnly(httptransport.NewServer(
		endpoints.CreateReferenceWithTopic,
		model.DecodeReferenceWithTopicReq,
		model.EncodeResponse,
		options...,
	)))

	r.Methods("POST").Path("/").Handler(httptransport.NewServer(
		endpoints.CreateReference,
		model.DecodeReferenceReq,
		model.EncodeResponse,
		options...,
	))

	return r
}

func adminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if true {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// } it will be process authentication
		ctx := r.Context()                               // get context background from request
		child := context.WithValue(ctx, "auth", "dhuki") // making child of parent context with value inside it
		req := r.WithContext(child)                      // bind ctx with request
		next.ServeHTTP(w, req)
	})
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
