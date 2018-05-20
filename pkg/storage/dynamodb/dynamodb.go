package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/rodrigodiez/zorro/pkg/storage"
)

type dynamodbStorage struct {
	svc         dynamodbiface.DynamoDBAPI
	keysTable   string
	valuesTable string
}

type item struct {
	ID   string
	Data string
}

type key struct {
	ID string
}

func (d *dynamodbStorage) LoadOrStore(key string, value string) (actualValue string, loaded bool) {

	_, err := d.svc.PutItem(newPutItemInput(key, value, d.keysTable))

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				output, _ := d.svc.GetItem(newGetItemInput(key, d.keysTable))
				actual := &item{}
				dynamodbattribute.UnmarshalMap(output.Item, actual)

				return actual.Data, true
			}
		}
	}

	d.svc.PutItem(newPutItemInput(value, key, d.valuesTable))

	return value, false
}

func (d *dynamodbStorage) Resolve(value string) (key string, ok bool) {
	output, _ := d.svc.GetItem(newGetItemInput(value, d.valuesTable))

	if len(output.Item) == 0 {
		return "", false
	}

	actual := &item{}
	dynamodbattribute.UnmarshalMap(output.Item, actual)

	return actual.Data, true
}

// New creates a new Storage persisted in AWS DynamoDB.
func New(svc dynamodbiface.DynamoDBAPI, keysTable string, valuesTable string) storage.Storage {
	return &dynamodbStorage{
		svc:         svc,
		keysTable:   keysTable,
		valuesTable: valuesTable,
	}
}

func newPutItemInput(id string, data string, table string) *dynamodb.PutItemInput {
	item := &item{ID: id, Data: data}
	dynamodbItem, _ := dynamodbattribute.MarshalMap(item)

	return &dynamodb.PutItemInput{
		Item:                dynamodbItem,
		TableName:           aws.String(table),
		ConditionExpression: aws.String("attribute_not_exists(ID)"),
	}
}

func newGetItemInput(id string, table string) *dynamodb.GetItemInput {

	dynamodbKey, _ := dynamodbattribute.MarshalMap(&key{ID: *aws.String(id)})

	return &dynamodb.GetItemInput{
		TableName:      aws.String(table),
		Key:            dynamodbKey,
		ConsistentRead: aws.Bool(true),
	}
}
