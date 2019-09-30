package services

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Theme struct {
	CorporateID int         `json:"corporateId"`
	Name        string      `json:"name"`
	Data        interface{} `json:"data"`
}

func AddThemeSchema(tableName string, corporateId int, name string, data interface{}) (resp interface{}, err error) {
	client := getClient()
	theme := Theme{
		CorporateID: corporateId,
		Name:        name,
		Data:        data,
	}

	fmt.Println(theme)
	// Map theme struct to data type descriptors
	attributes, err := dynamodbattribute.MarshalMap(theme)
	fmt.Println(attributes)

	// Create the PutItemInput struct to send the request
	input := &dynamodb.PutItemInput{
		Item:      attributes,
		TableName: aws.String(tableName),
	}

	var newTheme interface{}
	// Perform request to add the theme
	newTheme, err = client.PutItem(input)
	if err != nil {
		return nil, err
	}
	fmt.Println(newTheme)
	return newTheme, nil

}
