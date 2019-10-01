package server

import (
	"context"

	"github.com/StevenRojas/go-dynamodb-demo/services"
)

// Service interface
type Service interface {
	CreateSchema(ctx context.Context, request schemaRequest) (interface{}, error)
	GetSchema(ctx context.Context, request schemaDescriptionRequest) (interface{}, error)
	AddTheme(ctx context.Context, request addThemeRequest) (interface{}, error)
	ListTheme(ctx context.Context, request listThemeRequest) (interface{}, error)
	PatchTheme(ctx context.Context, request patchThemeRequest) (interface{}, error)
	QueryTheme(ctx context.Context, request queryThemeRequest) (interface{}, error)
	GetTheme(ctx context.Context, request getThemeRequest) (interface{}, error)
}

type dynamoDbService struct{}

// GetServiceInstance Get new instance of dynamoDB service
func GetServiceInstance() Service {
	return dynamoDbService{}
}

// CreateSchema Create a new schema
func (dynamoDbService) CreateSchema(ctx context.Context, request schemaRequest) (interface{}, error) {
	resp, err := services.CreateSchema(request.TableName)
	if err != nil {
		return "", err
	}
	return resp, nil
}

// GetSchema Get schema information
func (dynamoDbService) GetSchema(ctx context.Context, request schemaDescriptionRequest) (interface{}, error) {
	resp, err := services.GetSchema(request.TableName)
	if err != nil {
		return "", err
	}
	return resp, nil
}

// AddTheme add a new theme
func (dynamoDbService) AddTheme(ctx context.Context, request addThemeRequest) (interface{}, error) {
	resp, err := services.AddTheme(
		request.TableName,
		request.PartitionKey,
		request.SortKey,
		request.Data,
	)
	if err != nil {
		return "", err
	}
	return resp, nil
}

// ListTheme List of themes
func (dynamoDbService) ListTheme(ctx context.Context, request listThemeRequest) (interface{}, error) {
	resp, err := services.ListTheme(request.TableName)
	if err != nil {
		return "", err
	}
	return resp, nil
}

// PatchTheme Modify a theme
func (dynamoDbService) PatchTheme(ctx context.Context, request patchThemeRequest) (interface{}, error) {
	resp, err := services.PatchTheme(
		request.TableName,
		request.PartitionKey,
		request.SortKey,
		request.Data,
	)
	if err != nil {
		return "", err
	}
	return resp, nil
}

// QueryTheme Query a theme
func (dynamoDbService) QueryTheme(ctx context.Context, request queryThemeRequest) (interface{}, error) {
	resp, err := services.QueryTheme(request.TableName, request.Query)
	if err != nil {
		return "", err
	}
	return resp, nil
}

// GetTheme Get a theme
func (dynamoDbService) GetTheme(ctx context.Context, request getThemeRequest) (interface{}, error) {
	resp, err := services.GetTheme(request.TableName, request.Corp, request.Name)
	if err != nil {
		return "", err
	}
	return resp, nil
}
