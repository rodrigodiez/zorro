package dynamodb

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/rodrigodiez/zorro/pkg/storage"
)

type item struct {
	ID   string
	Data string
}

type key struct {
	ID string
}

type dynamodbStorage struct {
	svc         dynamodbiface.DynamoDBAPI
	keysTable   string
	valuesTable string
	metrics     *storage.Metrics
}

// New creates a new Storage persisted in AWS DynamoDB.
func New(svc dynamodbiface.DynamoDBAPI, keysTable string, valuesTable string) storage.Storage {
	return &dynamodbStorage{
		svc:         svc,
		keysTable:   keysTable,
		valuesTable: valuesTable,
	}
}

func (sto *dynamodbStorage) LoadOrStore(key string, value string) (string, error) {

	_, err := sto.svc.PutItem(newPutItemInput(key, value, sto.keysTable))

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				output, _ := sto.svc.GetItem(newGetItemInput(key, sto.keysTable))
				actual := &item{}
				dynamodbattribute.UnmarshalMap(output.Item, actual)

				sto.incrLoadOps()
				return actual.Data, nil
			}
		}

		return "", err
	}

	sto.svc.PutItem(newPutItemInput(value, key, sto.valuesTable))

	sto.incrStoreOps()
	return value, nil
}

func (sto *dynamodbStorage) Resolve(value string) (string, error) {
	output, err := sto.svc.GetItem(newGetItemInput(value, sto.valuesTable))
	sto.incrResolveOps()

	if err != nil {
		return "", err
	}

	if len(output.Item) == 0 {
		return "", errors.New("Key does not exist")
	}

	actual := &item{}
	dynamodbattribute.UnmarshalMap(output.Item, actual)

	return actual.Data, nil
}

// Close is noop
func (sto *dynamodbStorage) Close() {
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

func (sto *dynamodbStorage) WithMetrics(metrics *storage.Metrics) storage.Storage {
	sto.metrics = metrics

	return sto
}

func (sto *dynamodbStorage) incrStoreOps() {
	if sto.metrics != nil && sto.metrics.StoreOps != nil {
		sto.metrics.StoreOps.Add(int64(1))
	}
}

func (sto *dynamodbStorage) incrLoadOps() {
	if sto.metrics != nil && sto.metrics.LoadOps != nil {
		sto.metrics.LoadOps.Add(int64(1))
	}
}

func (sto *dynamodbStorage) incrResolveOps() {
	if sto.metrics != nil && sto.metrics.ResolveOps != nil {
		sto.metrics.ResolveOps.Add(int64(1))
	}
}
