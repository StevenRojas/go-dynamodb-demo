package server

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer(ctx context.Context, endpoints Endpoints) http.Handler {
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

	return r
}
