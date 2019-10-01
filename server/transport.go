package server

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type schemaRequest struct {
	TableName string `json:"tableName"`
}

type schemaDescriptionRequest struct {
	TableName string `string:"tableName"`
}

type addThemeRequest struct {
	TableName    string      `json:"tableName"`
	PartitionKey int         `json:"partitionKey"`
	SortKey      string      `json:"sortKey"`
	Data         interface{} `json:"data"`
}

type listThemeRequest struct {
	TableName string `string:"tableName"`
}

type patchThemeRequest struct {
	TableName    string      `json:"tableName"`
	PartitionKey int         `json:"partitionKey"`
	SortKey      string      `json:"sortKey"`
	Data         interface{} `json:"data"`
}

type queryThemeRequest struct {
	TableName string `string:"tableName"`
	Query     string `string:"query"`
}

type getThemeRequest struct {
	TableName string `string:"tableName"`
	Corp      int    `int:"corp"`
	Name      string `string:"name"`
}

type schemaResponse struct {
	Result interface{} `json:"response"`
}

type themeResponse struct {
	Result interface{} `json:"theme"`
}

func decodeSchemaRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req schemaRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeSchemaDescriptionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req schemaDescriptionRequest
	params := mux.Vars(r)
	req.TableName = params["tableName"]
	// TODO: Validate params
	return req, nil
}

func decodeAddThemeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req addThemeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeListThemeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req listThemeRequest
	params := mux.Vars(r)
	req.TableName = params["tableName"]
	return req, nil
}

func decodePatchThemeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req patchThemeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeQueryThemeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req queryThemeRequest
	params := mux.Vars(r)
	req.TableName = params["tableName"]
	req.Query = params["query"]
	return req, nil
}

func decodeGetThemeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getThemeRequest
	params := mux.Vars(r)
	req.TableName = params["tableName"]
	req.Corp, _ = strconv.Atoi(params["corp"])
	req.Name = params["name"]
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
