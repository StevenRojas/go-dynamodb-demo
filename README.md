# go-dynamodb-demo
This is demo project about Go, AWS SDK and DynamoDB. An explanation of the main concepts are at [DynamoDB Concepts](DynamoDB.md)

# Installation
$ `go get -u github.com/aws/aws-sdk-go/...`

$ `go get -u github.com/gorilla/mux`

# Execution
For local instance of DynamoDB with Docker (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.Docker.html)

$ `docker run -p 8000:8000 amazon/dynamodb-local`

Then run Go server (it is using Go Kit with Gorilla Mux)

$ `go run cmd/main.go`

Endpoints are at the [Postman collection](dynamo-demo.postman_collection.json)
