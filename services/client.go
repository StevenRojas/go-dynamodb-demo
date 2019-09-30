package services

import (
	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func getClient() *dynamodb.DynamoDB {
	// Define the config struct
	config := &aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String("http://localhost:8000"),
	}

	// Craete a new session
	sess := session.Must(session.NewSession(config))
	// Create client for the session
	client := dynamodb.New(sess)
	return client
}
