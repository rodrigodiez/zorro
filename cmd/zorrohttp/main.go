package main

import (
	"flag"
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/rodrigodiez/zorro"
)

func main() {
	var port int
	var storageDriver string
	var storagePath string
	var z zorro.Zorro

	flag.IntVar(&port, "port", 8080, "http port")
	flag.StringVar(&storageDriver, "storage-driver", "memory", "storage driver (memory, boltdb)")
	flag.StringVar(&storagePath, "storage-path", "", "storage path, mandatory when storage driver is boltdb")
	flag.Parse()

	switch storageDriver {
	case "memory":
		z = zorro.New(zorro.NewUUIDv4Generator(), zorro.NewInMemoryStorage())
	case "boltdb":
		if storagePath == "" {
			flag.Usage()
			os.Exit(-1)
		}

		storage, err := zorro.NewBoltDBStorage(storagePath)

		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		defer storage.Close()

		z = zorro.New(zorro.NewUUIDv4Generator(), storage)

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
