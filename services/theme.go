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

func AddTheme(tableName string, corporateId int, name string, data interface{}) (resp interface{}, err error) {
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

func ListTheme(tableName string) (resp interface{}, err error) {
	client := getClient()
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	result, err := client.Scan(params)
	if err != nil {
		return nil, err
	}
	themes := []Theme{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &themes)
	if err != nil {
		return nil, err
	}
	return themes, nil
}
