package main

import (
	"expvar"
	"flag"
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/rodrigodiez/zorro/lib/cli"
	"github.com/rodrigodiez/zorro/pkg/generator"
	"github.com/rodrigodiez/zorro/pkg/generator/uuid"
	"github.com/rodrigodiez/zorro/pkg/service"
	"github.com/rodrigodiez/zorro/pkg/storage"
)

func main() {
	var (
		zorro     service.Zorro
		generator generator.Generator
		sto       storage.Storage
		err       error
	)

	options := cli.GetOptions()

	flag.Parse()

	if *options.Help {
		flag.Usage()
		os.Exit(0)
	}

	sto, err = cli.GetStorageForOptions(options)

	if err != nil {
		flag.Usage()
		log.Fatal(err)
		os.Exit(-1)
	}

	defer sto.Close()

	generator = uuid.NewV4()
	zorro = service.New(generator, sto)

	if *options.Debug {
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
		port:   *options.Port,
	}

	server.serve(*options.Debug)
}
