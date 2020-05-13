package router

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	transport "github.com/website-pribadi/pkg/topic/interface/transport"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/create-topic").Handler(httptransport.NewServer(
		endpoints.CreateTopic,
		transport.DecodeTopicReq,
		transport.EncodeResponse,
	))

	// r.Methods("GET").Path("/{id}").Handler(httptransport.NewServer(

	// ))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
