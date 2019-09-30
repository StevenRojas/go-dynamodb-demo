package services

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetSchema(tableName string) (resp interface{}, err error) {
	client := getClient()
	request := &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	}
	result, err := client.DescribeTable(request)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	//fmt.Println(result)
	return result, nil
}

func CreateSchema(tableName string) (resp interface{}, err error) {
	client := getClient()
	schema := getTableSchema(tableName)

	result, err := client.CreateTable(schema)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return result, nil

}

func getTableSchema(tableName string) *dynamodb.CreateTableInput {
	schema := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("corporateId"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("name"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("corporateId"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("name"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}
	return schema
}
