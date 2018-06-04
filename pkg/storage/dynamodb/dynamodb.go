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

type dynamodbStorage struct {
	svc         dynamodbiface.DynamoDBAPI
	keysTable   string
	valuesTable string
	metrics     *storage.Metrics
}

type item struct {
	ID   string
	Data string
}

type key struct {
	ID string
}

func (d *dynamodbStorage) LoadOrStore(key string, value string) (string, error) {

	_, err := d.svc.PutItem(newPutItemInput(key, value, d.keysTable))

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				output, _ := d.svc.GetItem(newGetItemInput(key, d.keysTable))
				actual := &item{}
				dynamodbattribute.UnmarshalMap(output.Item, actual)

				d.incrLoadOps()
				return actual.Data, nil
			}
		}

		return "", err
	}

	d.svc.PutItem(newPutItemInput(value, key, d.valuesTable))

	d.incrStoreOps()
	return value, nil
}

func (d *dynamodbStorage) Resolve(value string) (string, error) {
	output, err := d.svc.GetItem(newGetItemInput(value, d.valuesTable))
	d.incrResolveOps()

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

func (d *dynamodbStorage) Close() {
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

func (d *dynamodbStorage) WithMetrics(metrics *storage.Metrics) storage.Storage {
	d.metrics = metrics

	return d
}

func (d *dynamodbStorage) incrStoreOps() {
	if d.metrics != nil && d.metrics.StoreOps != nil {
		d.metrics.StoreOps.Add(int64(1))
	}
}

func (d *dynamodbStorage) incrLoadOps() {
	if d.metrics != nil && d.metrics.LoadOps != nil {
		d.metrics.LoadOps.Add(int64(1))
	}
}

func (d *dynamodbStorage) incrResolveOps() {
	if d.metrics != nil && d.metrics.ResolveOps != nil {
		d.metrics.ResolveOps.Add(int64(1))
	}
}
