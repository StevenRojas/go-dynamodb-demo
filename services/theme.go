package services

import (
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// ThemeKey struct. PartitionKey and SortKey
type ThemeKey struct {
	CorporateID int    `json:"corporateId"`
	Name        string `json:"name"`
}

// Theme struct PartitionKey, SortKey and Data
type Theme struct {
	CorporateID int         `json:"corporateId"`
	Name        string      `json:"name"`
	Data        interface{} `json:"data"`
}

// Settings struct (a theme node)
type Settings struct {
	MessageText  string `json:"messageText"`
	TimeoutAfter int    `json:"timeoutAfter"`
}

// SettingsData struct that contains the settings node
type SettingsData struct {
	Settings Settings `json:"settings"`
}

// AddTheme add a new theme to dynamoDB
func AddTheme(tableName string, corporateID int, name string, data interface{}) (resp interface{}, err error) {
	client := getClient()
	theme := Theme{
		CorporateID: corporateID,
		Name:        name,
		Data:        data,
	}

	// Map theme struct to data type descriptors
	attributes, err := dynamodbattribute.MarshalMap(theme)

	// Create the PutItemInput struct to send the request
	input := &dynamodb.PutItemInput{
		Item:      attributes,
		TableName: aws.String(tableName),
	}

	// Perform request to add the theme
	_, err = client.PutItem(input)
	if err != nil {
		return nil, err
	}
	themeKey := ThemeKey{
		CorporateID: corporateID,
		Name:        name,
	}
	return themeKey, nil

}

// ListTheme list of themes stored at dynamodDB
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

// GetTheme get a theme stored at dynamodDB
func GetTheme(tableName string, corporateID int, name string) (resp interface{}, err error) {
	client := getClient()
	themeKey := ThemeKey{
		CorporateID: corporateID,
		Name:        name,
	}
	key, err := dynamodbattribute.MarshalMap(themeKey)
	input := &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(tableName),
	}

	result, err := client.GetItem(input)
	if err != nil {
		return nil, err
	}

	theme := Theme{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &theme)
	if err != nil {
		return nil, err
	}
	return theme, nil
}

// QueryTheme query for themes stored at dynamodDB
func QueryTheme(tableName string, query string) (resp interface{}, err error) {
	client := getClient()
	conditions := map[string]*dynamodb.Condition{
		"name": {
			ComparisonOperator: aws.String("EQ"),
			AttributeValueList: []*dynamodb.AttributeValue{
				{
					S: aws.String(query),
				},
			},
		},
	}

	// TODO: Create index mediaType
	queryInput := &dynamodb.QueryInput{
		TableName:     aws.String(tableName),
		IndexName:     aws.String("mediaType"),
		KeyConditions: conditions,
	}
	result, err := client.Query(queryInput)
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

// PatchTheme update a theme stored at dynamodDB
func PatchTheme(tableName string, corporateID int, name string, data interface{}) (resp interface{}, err error) {
	client := getClient()
	themeKey := ThemeKey{
		CorporateID: corporateID,
		Name:        name,
	}

	// Get a key for the theme that will be updated
	key, err := dynamodbattribute.MarshalMap(themeKey)
	if err != nil {
		return nil, err
	}

	// Decode the data field to define what will be updated
	body := data.(map[string]interface{})

	// Check if "settings" field is set
	if _, ok := body["settings"]; ok {
		result, err := updateSettings(tableName, key, body, client)
		fmt.Println(reflect.TypeOf(result))
		fmt.Println(result)
		fmt.Println(err)
	}

	return nil, err
}

func updateSettings(tableName string, key map[string]*dynamodb.AttributeValue,
	body map[string]interface{}, client *dynamodb.DynamoDB) (resp interface{}, err error) {

	sett := body["settings"].(map[string]interface{})
	settings := Settings{
		MessageText:  sett["messageText"].(string),
		TimeoutAfter: int(sett["timeoutAfter"].(float64)),
	}
	update, err := dynamodbattribute.MarshalMap(SettingsData{
		Settings: settings,
	})
	fmt.Println(update)
	// FIXME: Something is wrong with the update values
	input := &dynamodb.UpdateItemInput{
		Key:                       key,
		TableName:                 aws.String(tableName),
		UpdateExpression:          aws.String("SET settings.messageText = :m, settings.timeoutAfter = :t"),
		ExpressionAttributeValues: update,
		ReturnValues:              aws.String("UPDATED_NEW"),
	}
	result, err := client.UpdateItem(input)
	fmt.Println(result, err)
	// TODO Parse result and return a response
	return nil, nil
}
