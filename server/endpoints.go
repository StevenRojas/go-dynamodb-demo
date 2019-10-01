package server

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints struct with the list of all available endpoints
type Endpoints struct {
	SchemaEndpoint            endpoint.Endpoint
	SchemaDescriptionEndpoint endpoint.Endpoint
	AddThemeEndpoint          endpoint.Endpoint
	ListThemeEndpoint         endpoint.Endpoint
	PatchThemeEndpoint        endpoint.Endpoint
	QueryThemeEndpoint        endpoint.Endpoint
	GetThemeEndpoint          endpoint.Endpoint
}

// MakeSchemaEndpoint create schema endpoint
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

// MakeSchemaDescriptionEndpoint get schema endpoint
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

//MakeAddThemeEndpoint add a new theme
func MakeAddThemeEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addThemeRequest)
		s, err := srv.AddTheme(ctx, req)

		if err != nil {
			return themeResponse{s}, err
		}
		return themeResponse{s}, nil
	}
}

// MakeListThemeEndpoint get a list of themes
func MakeListThemeEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(listThemeRequest)
		s, err := srv.ListTheme(ctx, req)

		if err != nil {
			return themeResponse{s}, err
		}
		return themeResponse{s}, nil
	}
}

// MakePatchThemeEndpoint modify a theme
func MakePatchThemeEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(patchThemeRequest)
		s, err := srv.PatchTheme(ctx, req)

		if err != nil {
			return themeResponse{s}, err
		}
		return themeResponse{s}, nil
	}
}

// MakeQueryThemeEndpoint perform a query over themes
func MakeQueryThemeEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(queryThemeRequest)
		s, err := srv.QueryTheme(ctx, req)

		if err != nil {
			return themeResponse{s}, err
		}
		return themeResponse{s}, nil
	}
}

// MakeGetThemeEndpoint get a theme
func MakeGetThemeEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getThemeRequest)
		s, err := srv.GetTheme(ctx, req)

		if err != nil {
			return themeResponse{s}, err
		}
		return themeResponse{s}, nil
	}
}
