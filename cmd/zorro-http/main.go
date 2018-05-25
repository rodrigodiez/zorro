package main

import (
	"expvar"
	"flag"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	awsDynamodb "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gorilla/mux"
	"github.com/rodrigodiez/zorro/pkg/generator"
	"github.com/rodrigodiez/zorro/pkg/generator/uuid"
	"github.com/rodrigodiez/zorro/pkg/service"
	"github.com/rodrigodiez/zorro/pkg/storage"
	"github.com/rodrigodiez/zorro/pkg/storage/boltdb"
	"github.com/rodrigodiez/zorro/pkg/storage/dynamodb"
	"github.com/rodrigodiez/zorro/pkg/storage/memory"
)

func main() {
	var (
		port                int
		storageDriver       string
		storagePath         string
		dynamodbKeysTable   string
		dynamodbValuesTable string
		awsRegion           string
		help                bool
		debug               bool
		zorro               service.Zorro
		generator           generator.Generator
		sto                 storage.Storage
		err                 error
	)

	flag.IntVar(&port, "port", 8080, "http port")
	flag.StringVar(&storageDriver, "storage-driver", "memory", "storage driver (memory, boltdb, dynamodb)")
	flag.StringVar(&storagePath, "storage-path", "", "storage path, mandatory for boltdb as storage driver")
	flag.StringVar(&dynamodbKeysTable, "dynamodb-keys-table", "", "dynamodb table where to store the keys for resolving values, mandatory for dynamodb as storage driver")
	flag.StringVar(&dynamodbValuesTable, "dynamodb-values-table", "", "dynamodb table where to store the values to resolve keys, mandatory for dynamodb as storage driver")
	flag.StringVar(&awsRegion, "aws-region", "", "AWS region (ie: eu-west-1), mandatory for dynamodb as storage driver")
	flag.BoolVar(&debug, "debug", false, "serves metrics on /debug/vars")
	flag.BoolVar(&help, "help", false, "prints help")

	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	switch storageDriver {
	case "memory":
		sto = memory.New()
	case "boltdb":
		if storagePath == "" {
			flag.Usage()
			os.Exit(-1)
		}

		sto, err = boltdb.New(storagePath)

		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
	case "dynamodb":
		if dynamodbKeysTable == "" || dynamodbValuesTable == "" || awsRegion == "" {
			flag.Usage()
			os.Exit(-1)
		}

		sess, _ := session.NewSession(&aws.Config{Region: &awsRegion})
		sto = dynamodb.New(awsDynamodb.New(sess), dynamodbKeysTable, dynamodbValuesTable)
	default:
		flag.Usage()
		os.Exit(-1)
	}

	defer sto.Close()

	generator = uuid.NewV4()
	zorro = service.New(generator, sto)

	if debug {
		var (
			maskOps    expvar.Int
			unmaskOps  expvar.Int
			storeOps   expvar.Int
			loadOps    expvar.Int
			resolveOps expvar.Int
		)

		metricsMap := expvar.NewMap("zorro")
		metricsMap.Set("maskOps", &maskOps)
		metricsMap.Set("unmaskOps", &unmaskOps)
		metricsMap.Set("storeOps", &storeOps)
		metricsMap.Set("loadOps", &loadOps)
		metricsMap.Set("resolveOps", &resolveOps)

		zorro.WithMetrics(&service.Metrics{MaskOps: &maskOps, UnmaskOps: &unmaskOps})
		sto.WithMetrics(&storage.Metrics{StoreOps: &storeOps, LoadOps: &loadOps, ResolveOps: &resolveOps})
	}

	server := &server{
		router: mux.NewRouter(),
		zorro:  zorro,
		port:   port,
	}

	server.serve(debug)
}
