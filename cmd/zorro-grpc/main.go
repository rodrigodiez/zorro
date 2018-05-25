package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/rodrigodiez/zorro/lib/cli"
	"github.com/rodrigodiez/zorro/pkg/generator"
	"github.com/rodrigodiez/zorro/pkg/generator/uuid"
	"github.com/rodrigodiez/zorro/pkg/protobuf"
	"github.com/rodrigodiez/zorro/pkg/service"
	"github.com/rodrigodiez/zorro/pkg/storage"
	"google.golang.org/grpc"
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
		log.Fatal(err)
		flag.Usage()
		os.Exit(-1)
	}

	defer sto.Close()

	generator = uuid.NewV4()
	zorro = service.New(generator, sto)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *options.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		os.Exit(-1)
	}
	grpcServer := grpc.NewServer()
	protobuf.RegisterZorroServer(grpcServer, &server{zorro: zorro, storage: sto, generator: generator})
	log.Printf("Listening for connections on :%d\n", *options.Port)
	grpcServer.Serve(lis)
}
