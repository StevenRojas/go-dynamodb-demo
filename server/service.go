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
	AddThemeSchema(ctx context.Context, request addThemeRequest) (interface{}, error)
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
	fmt.Println(resp)
	return resp, nil
}

func (dynamoDbService) GetSchema(ctx context.Context, request schemaDescriptionRequest) (interface{}, error) {
	resp, err := services.GetSchema(request.TableName)
	if err != nil {
		return "", err
	}
	fmt.Println(resp)
	return resp, nil
}

func (dynamoDbService) AddThemeSchema(ctx context.Context, request addThemeRequest) (interface{}, error) {
	resp, err := services.AddThemeSchema(
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
