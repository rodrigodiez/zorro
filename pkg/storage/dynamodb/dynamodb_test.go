package dynamodb

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	dynamodbapiMocks "github.com/rodrigodiez/zorro/lib/mocks/dynamodbapi"
	metricsMocks "github.com/rodrigodiez/zorro/lib/mocks/metrics"
	"github.com/rodrigodiez/zorro/pkg/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewImplementsStorage(t *testing.T) {
	var _ storage.Storage = New(&dynamodbapiMocks.DynamoDBAPI{}, "keysTable", "valuesTable")
}

func TestLoadOrStoreReturnsValueAndNilIfKeyDoesNotExist(t *testing.T) {
	t.Parallel()

	svc := &dynamodbapiMocks.DynamoDBAPI{}
	storage := New(svc, "keysTable", "valuesTable")

	svc.On("PutItem", mock.MatchedBy(func(input *dynamodb.PutItemInput) bool {
		return *input.TableName == "keysTable" && *input.Item["ID"].S == "foo" && *input.Item["Data"].S == "bar"
	})).Return(nil, nil).Once()

	svc.On("PutItem", mock.MatchedBy(func(input *dynamodb.PutItemInput) bool {
		return *input.TableName == "valuesTable" && *input.Item["ID"].S == "bar" && *input.Item["Data"].S == "foo"
	})).Return(nil, nil).Once()

	value, err := storage.LoadOrStore("foo", "bar")

	assert.Equal(t, "bar", value)
	assert.Nil(t, err)

	svc.AssertExpectations(t)
}

func TestLoadOrStoreReturnsActualValueAndNilIfKeyExists(t *testing.T) {
	t.Parallel()

	svc := &dynamodbapiMocks.DynamoDBAPI{}
	storage := New(svc, "keysTable", "valuesTable")

	svc.On("PutItem", mock.MatchedBy(func(input *dynamodb.PutItemInput) bool {
		return *input.TableName == "keysTable" && *input.ConditionExpression == "attribute_not_exists(ID)"
	})).Return(nil, awserr.New(dynamodb.ErrCodeConditionalCheckFailedException, "", nil)).Once()

	svc.On("GetItem", mock.MatchedBy(func(input *dynamodb.GetItemInput) bool {
		return *input.TableName == "keysTable" && *input.Key["ID"].S == "foo" && *input.ConsistentRead
	})).Return(newPutItemOutput("foo", "baz"), nil).Once()

	value, err := storage.LoadOrStore("foo", "bar")

	assert.Equal(t, "baz", value)
	assert.Nil(t, err)

	svc.AssertExpectations(t)
}

func TestLoadOrStoreReturnsEmptyStringAndErrorIfStorageFails(t *testing.T) {
	t.Parallel()

	svc := &dynamodbapiMocks.DynamoDBAPI{}
	storage := New(svc, "keysTable", "valuesTable")

	svc.On("PutItem", mock.MatchedBy(func(input *dynamodb.PutItemInput) bool {
		return *input.TableName == "keysTable" && *input.ConditionExpression == "attribute_not_exists(ID)"
	})).Return(nil, awserr.New(dynamodb.ErrCodeInternalServerError, "", nil)).Once()

	value, err := storage.LoadOrStore("foo", "bar")

	assert.Equal(t, "", value)
	assert.NotNil(t, err)

	svc.AssertExpectations(t)
}

func TestResolveReturnsKeyAndNilIfExists(t *testing.T) {
	t.Parallel()

	svc := &dynamodbapiMocks.DynamoDBAPI{}
	storage := New(svc, "keysTable", "valuesTable")

	svc.On("GetItem", mock.MatchedBy(func(input *dynamodb.GetItemInput) bool {
		return *input.TableName == "valuesTable" && *input.Key["ID"].S == "bar" && *input.ConsistentRead
	})).Return(newPutItemOutput("bar", "foo"), nil).Once()

	key, err := storage.Resolve("bar")

	assert.Equal(t, "foo", key)
	assert.Nil(t, err)

	svc.AssertExpectations(t)
}
func TestResolveReturnsEmptyAndErrIfNotExists(t *testing.T) {
	t.Parallel()

	svc := &dynamodbapiMocks.DynamoDBAPI{}
	storage := New(svc, "keysTable", "valuesTable")

	svc.On("GetItem", mock.MatchedBy(func(input *dynamodb.GetItemInput) bool {
		return *input.TableName == "valuesTable" && *input.Key["ID"].S == "bar" && *input.ConsistentRead
	})).Return(&dynamodb.GetItemOutput{Item: make(map[string]*dynamodb.AttributeValue)}, nil).Once()

	key, err := storage.Resolve("bar")

	assert.Equal(t, "", key)
	assert.NotNil(t, err)

	svc.AssertExpectations(t)
}

func TestResolveReturnsEmptyAndErrIfStorageFails(t *testing.T) {
	t.Parallel()

	svc := &dynamodbapiMocks.DynamoDBAPI{}
	storage := New(svc, "keysTable", "valuesTable")

	svc.On("GetItem", mock.MatchedBy(func(input *dynamodb.GetItemInput) bool {
		return *input.TableName == "valuesTable" && *input.Key["ID"].S == "bar" && *input.ConsistentRead
	})).Return(nil, awserr.New(dynamodb.ErrCodeInternalServerError, "", nil)).Once()

	key, err := storage.Resolve("bar")

	assert.Equal(t, "", key)
	assert.NotNil(t, err)

	svc.AssertExpectations(t)
}

func TestLoadOrStoreIncrementsStoreOpsIfKeyDoesNotExist(t *testing.T) {
	t.Parallel()

	counter := &metricsMocks.IntCounter{}
	counter.On("Add", int64(1))
	svc := &dynamodbapiMocks.DynamoDBAPI{}

	storage := New(svc, "keysTable", "valuesTable").WithMetrics(&storage.Metrics{StoreOps: counter})

	svc.On("PutItem", mock.MatchedBy(func(input *dynamodb.PutItemInput) bool {
		return *input.TableName == "keysTable" && *input.Item["ID"].S == "foo" && *input.Item["Data"].S == "bar"
	})).Return(nil, nil).Once()

	svc.On("PutItem", mock.MatchedBy(func(input *dynamodb.PutItemInput) bool {
		return *input.TableName == "valuesTable" && *input.Item["ID"].S == "bar" && *input.Item["Data"].S == "foo"
	})).Return(nil, nil).Once()

	storage.LoadOrStore("foo", "bar")

	counter.AssertCalled(t, "Add", int64(1))
}

func TestLoadOrStoreIncrementsLoadOpsIfKeyExists(t *testing.T) {
	t.Parallel()

	counter := &metricsMocks.IntCounter{}
	counter.On("Add", int64(1))
	svc := &dynamodbapiMocks.DynamoDBAPI{}

	storage := New(svc, "keysTable", "valuesTable").WithMetrics(&storage.Metrics{LoadOps: counter})

	svc.On("PutItem", mock.Anything).Return(nil, awserr.New(dynamodb.ErrCodeConditionalCheckFailedException, "", nil)).Maybe()
	svc.On("GetItem", mock.Anything).Return(newPutItemOutput("foo", "baz"), nil).Maybe()

	storage.LoadOrStore("foo", "bar")

	counter.AssertCalled(t, "Add", int64(1))
}

func TestResolveIncrementsResolveOps(t *testing.T) {
	t.Parallel()

	counter := &metricsMocks.IntCounter{}
	counter.On("Add", int64(1))

	svc := &dynamodbapiMocks.DynamoDBAPI{}
	storage := New(svc, "keysTable", "valuesTable").WithMetrics(&storage.Metrics{ResolveOps: counter})
	svc.On("GetItem", mock.Anything).Return(&dynamodb.GetItemOutput{Item: make(map[string]*dynamodb.AttributeValue)}, nil).Once()

	storage.Resolve("bar")

	counter.AssertCalled(t, "Add", int64(1))
}

func newPutItemOutput(id string, data string) *dynamodb.GetItemOutput {
	item := &item{ID: id, Data: data}
	dynamodbItem, _ := dynamodbattribute.MarshalMap(item)

	return &dynamodb.GetItemOutput{
		Item: dynamodbItem,
	}
}
