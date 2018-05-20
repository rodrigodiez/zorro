package main

import (
	"flag"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	awsDynamodb "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gorilla/mux"
	"github.com/rodrigodiez/zorro/pkg/generator/uuid"
	"github.com/rodrigodiez/zorro/pkg/service"
	"github.com/rodrigodiez/zorro/pkg/storage/boltdb"
	"github.com/rodrigodiez/zorro/pkg/storage/dynamodb"
	"github.com/rodrigodiez/zorro/pkg/storage/memory"
)

func main() {
	var port int
	var storageDriver string
	var storagePath string
	var dynamodbKeysTable string
	var dynamodbValuesTable string
	var awsRegion string
	var help bool
	var z service.Zorro

	flag.IntVar(&port, "port", 8080, "http port")
	flag.StringVar(&storageDriver, "storage-driver", "memory", "storage driver (memory, boltdb, dynamodb)")
	flag.StringVar(&storagePath, "storage-path", "", "storage path, mandatory for boltdb as storage driver")
	flag.StringVar(&dynamodbKeysTable, "dynamodb-keys-table", "", "dynamodb table where to store the keys for resolving values, mandatory for dynamodb as storage driver")
	flag.StringVar(&dynamodbValuesTable, "dynamodb-values-table", "", "dynamodb table where to store the values to resolve keys, mandatory for dynamodb as storage driver")
	flag.StringVar(&awsRegion, "aws-region", "", "AWS region (ie: eu-west-1), mandatory for dynamodb as storage driver")
	flag.BoolVar(&help, "help", false, "prints help")
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	switch storageDriver {
	case "memory":
		z = service.New(uuid.NewV4(), memory.New())
	case "boltdb":
		if storagePath == "" {
			flag.Usage()
			os.Exit(-1)
		}

		storage, err := boltdb.New(storagePath)

		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		defer storage.Close()

		z = service.New(uuid.NewV4(), storage)
	case "dynamodb":
		if dynamodbKeysTable == "" || dynamodbValuesTable == "" || awsRegion == "" {
			flag.Usage()
			os.Exit(-1)
		}

		sess, _ := session.NewSession(&aws.Config{Region: &awsRegion})
		storage := dynamodb.New(awsDynamodb.New(sess), dynamodbKeysTable, dynamodbValuesTable)
		z = service.New(uuid.NewV4(), storage)

	default:
		flag.Usage()
		os.Exit(-1)
	}

	a := &app{
		router: mux.NewRouter(),
		z:      z,
		port:   port,
	}

	a.run()
}
