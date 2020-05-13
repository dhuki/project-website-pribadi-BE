package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	Request struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	Response struct {
		ID          string `json:"id"`
		Name        string `json:"Name"`
		Description string `json:"description"`
	}
)

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func DecodeTopicReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeTopicGetById(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		return nil, errors.New("id is empty")
	}

	req := Request{
		ID: id,
	}
	return req, nil
}
