package cli

import (
	"context"
	"errors"
	"flag"
	"log"

	goDatastore "cloud.google.com/go/datastore"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	awsDynamodb "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/rodrigodiez/zorro/pkg/storage"
	"github.com/rodrigodiez/zorro/pkg/storage/boltdb"
	"github.com/rodrigodiez/zorro/pkg/storage/datastore"
	"github.com/rodrigodiez/zorro/pkg/storage/dynamodb"
	"github.com/rodrigodiez/zorro/pkg/storage/memory"
)

// Options is a type that holds pointers to cli flags shared by various servers (http, grpc...)
type Options struct {
	Port                *int
	StorageDriver       *string
	StoragePath         *string
	DynamoDBKeysTable   *string
	DynamoDBValuesTable *string
	AwsRegion           *string
	GcProjectID         *string
	DatastoreKeysKind   *string
	DatastoreValuesKind *string
	Debug               *bool
	Help                *bool
}

// GetOptions returns an Options struct based on cli flags. It has to be called BEFORE flag.Parse()
func GetOptions() *Options {

	return &Options{
		Port:                flag.Int("port", 8080, "http port"),
		StorageDriver:       flag.String("storage-driver", "memory", "storage driver (memory, boltdb, dynamodb, datastore)"),
		StoragePath:         flag.String("storage-path", "", "storage path, mandatory if storage driver is boltdb"),
		DynamoDBKeysTable:   flag.String("dynamodb-keys-table", "", "dynamodb table where to store keys, mandatory if storage driver is dynamodb"),
		DynamoDBValuesTable: flag.String("dynamodb-values-table", "", "dynamodb table where to store values, mandatory if storage driver is dynamodb"),
		AwsRegion:           flag.String("aws-region", "", "AWS region (ie: eu-west-1), mandatory if storage driver is dynamodb"),
		GcProjectID:         flag.String("gc-project-id", "", "Google Cloud project ID, mandatory if storage driver is datastore"),
		DatastoreKeysKind:   flag.String("datastore-keys-kind", "", "datastore kind where to store keys, mandatory if storage driver is datastore"),
		DatastoreValuesKind: flag.String("datastore-values-kind", "", "datastore kind where to store values, mandatory if storage driver is datastore"),
		Debug:               flag.Bool("debug", false, "serves metrics on /debug/vars"),
		Help:                flag.Bool("help", false, "prints help"),
	}
}

// GetStorageForOptions returns the appriate storage driver instance based on options
func GetStorageForOptions(options *Options) (storage.Storage, error) {
	var (
		sto storage.Storage
		err error
	)

	switch *options.StorageDriver {
	case "memory":
		sto = memory.New()
	case "boltdb":
		if *options.StoragePath == "" {
			return nil, errors.New("No storage path")
		}

		sto, err = boltdb.New(*options.StoragePath)

		if err != nil {
			return nil, err
		}
	case "dynamodb":
		if *options.DynamoDBKeysTable == "" || *options.DynamoDBValuesTable == "" || *options.AwsRegion == "" {
			return nil, errors.New("Some dynamodb options are missing")
		}

		sess, _ := session.NewSession(&aws.Config{Region: options.AwsRegion})
		sto = dynamodb.New(awsDynamodb.New(sess), *options.DynamoDBKeysTable, *options.DynamoDBValuesTable)
	case "datastore":
		if *options.DatastoreKeysKind == "" || *options.DatastoreValuesKind == "" || *options.GcProjectID == "" {
			return nil, errors.New("Some datastore options are missing")
		}

		client, err := goDatastore.NewClient(context.TODO(), *options.GcProjectID)
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}
		sto = datastore.New(datastore.NewTranslator(client), *options.DatastoreKeysKind, *options.DatastoreValuesKind)
	default:
		return nil, errors.New("Unknown storage driver")
	}

	return sto, nil
}
