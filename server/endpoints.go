package server

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	SchemaEndpoint            endpoint.Endpoint
	SchemaDescriptionEndpoint endpoint.Endpoint
	AddThemeEndpoint          endpoint.Endpoint
}

func MakeSchemaEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schemaRequest)
		s, err := srv.CreateSchema(ctx, req)

		if err != nil {
			return schemaResponse{s}, err
		}
		return schemaResponse{s}, nil
	}
}

func MakeSchemaDescriptionEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schemaDescriptionRequest)
		s, err := srv.GetSchema(ctx, req)

		if err != nil {
			return schemaResponse{s}, err
		}
		return schemaResponse{s}, nil
	}
}

func MakeAddThemeEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addThemeRequest)
		s, err := srv.AddThemeSchema(ctx, req)

		if err != nil {
			return themeResponse{s}, err
		}
		return themeResponse{s}, nil
	}
}
