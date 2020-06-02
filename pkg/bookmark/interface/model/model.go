package model

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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
	ReferenceTopicRequest struct {
		NameTopic        string   `json:"nameTopic"`
		DescriptionTopic string   `json:"descriptionTopic"`
		Links            []string `json:"links"`
	}

	ReferenceTopicResponse struct {
		ID               string            `json:"id"`
		NameTopic        string            `json:"name"`
		DescriptionTopic string            `json:"description"`
		Reference        ReferenceResponse `json:"reference"`
	}
)

func DecodeReferenceWithTopicReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req ReferenceTopicRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

type (
	ReferenceRequest struct {
		ID      string `json:"id"`
		TopicID string `json:"topicId"`
		Link    string `json:"link"`
	}

	ReferenceResponse struct {
		ID        string `json:"id"`
		TopicID   string `json:"topicId"`
		TopicName string `json:"topicName"`
		Link      string `json:"link"`
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

// func DecodeTopicGetById(ctx context.Context, r *http.Request) (interface{}, error) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	if id == "" {
// 		return nil, errors.New("id is empty")
// 	}

// 	req := TopicRequest{
// 		ID: id,
// 	}
// 	return req, nil
// }

// func DecodeTopicGetAll(ctx context.Context, r *http.Request) (interface{}, error) {
// 	return TopicRequest{}, nil
// }
