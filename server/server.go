package server

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer craetes a new HTTP server with the corresponding endpoints
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()

	r.Methods("GET").Path("/schema/{tableName}").Handler(httptransport.NewServer(
		endpoints.SchemaDescriptionEndpoint,
		decodeSchemaDescriptionRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/schema").Handler(httptransport.NewServer(
		endpoints.SchemaEndpoint,
		decodeSchemaRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/theme").Handler(httptransport.NewServer(
		endpoints.AddThemeEndpoint,
		decodeAddThemeRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/theme/{tableName}").Handler(httptransport.NewServer(
		endpoints.ListThemeEndpoint,
		decodeListThemeRequest,
		encodeResponse,
	))

	r.Methods("PATCH").Path("/theme").Handler(httptransport.NewServer(
		endpoints.PatchThemeEndpoint,
		decodePatchThemeRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/theme/{tableName}/query/{query}").Handler(httptransport.NewServer(
		endpoints.QueryThemeEndpoint,
		decodeQueryThemeRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/theme/{tableName}/corp/{corp}/name/{name}").Handler(httptransport.NewServer(
		endpoints.GetThemeEndpoint,
		decodeGetThemeRequest,
		encodeResponse,
	))

	return r
}
