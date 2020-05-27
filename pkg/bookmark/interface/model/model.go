package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type BaseResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	ErrCode int16       `json:"errCode"`
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	fmt.Println(response)
	return json.NewEncoder(w).Encode(response)
}

type (
	TopicRequest struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	TopicResponse struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}
)

func DecodeTopicReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req TopicRequest
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

	req := TopicRequest{
		ID: id,
	}
	return req, nil
}

func DecodeTopicGetAll(ctx context.Context, r *http.Request) (interface{}, error) {
	return TopicRequest{}, nil
}

type (
	ReferenceRequest struct {
		ID      string `json:"id"`
		TopicId string `json:"topicId"`
		Link    string `json:"link"`
	}
)

func DecodeReferenceReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req ReferenceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
