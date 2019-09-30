package server

import (
	"context"
	"fmt"

	"github.com/StevenRojas/go-dynamodb-demo/services"
)

// Service interface
type Service interface {
	CreateSchema(ctx context.Context, request schemaRequest) (interface{}, error)
	GetSchema(ctx context.Context, request schemaDescriptionRequest) (interface{}, error)
	AddTheme(ctx context.Context, request addThemeRequest) (interface{}, error)
	ListTheme(ctx context.Context, request listThemeRequest) (interface{}, error)
}

type dynamoDbService struct{}

func GetServiceInstance() Service {
	return dynamoDbService{}
}

func (dynamoDbService) CreateSchema(ctx context.Context, request schemaRequest) (interface{}, error) {
	resp, err := services.CreateSchema(request.TableName)
	if err != nil {
		return "", err
	}
	return resp, nil
}

func (dynamoDbService) GetSchema(ctx context.Context, request schemaDescriptionRequest) (interface{}, error) {
	resp, err := services.GetSchema(request.TableName)
	if err != nil {
		return "", err
	}
	return resp, nil
}

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
	fmt.Println(resp)
	return resp, nil
}

func (dynamoDbService) ListTheme(ctx context.Context, request listThemeRequest) (interface{}, error) {
	resp, err := services.ListTheme(request.TableName)
	if err != nil {
		return "", err
	}
	fmt.Println(resp)
	return resp, nil
}
