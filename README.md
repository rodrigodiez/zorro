# Zorro

[![build](	https://img.shields.io/travis/rodrigodiez/zorro.svg)](https://travis-ci.org/rodrigodiez/zorro)
[![Go Report Card](https://goreportcard.com/badge/github.com/rodrigodiez/zorro)](https://goreportcard.com/report/github.com/rodrigodiez/zorro)
[![](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/rodrigodiez/zorro)
[![](https://images.microbadger.com/badges/image/rodrigodiez/zorrohttp.svg)](https://microbadger.com/images/rodrigodiez/zorrohttp "Get your own image badge on microbadger.com")
[![MIT License](https://img.shields.io/github/license/rodrigodiez/zorro.svg)](https://github.com/rodrigodiez/zorro/blob/master/LICENSE.md)

Zorro allows developers to perform two way key/value lookups.

Zorro can be used with one of the provided servers or as a golang package for maximum flexibility.

[![gopher](https://github.com/egonelbre/gophers/raw/master/.thumb/vector/superhero/standing.png)](https://github.com/egonelbre/gophers)

by [@egonelbre](https://github.com/egonelbre/gophers)

---

> **Important**: Zorro is under heavy development at the moment and its usage in production is **not** recommended

## Use cases
- Services that want to protect their private IDs by translating them to public ones while keeping the ability to translate them back

## Running a Zorro

> At the moment only an http server is available

Easiest way to get your hands into Zorro is by running the docker image for the http server

```bash
# Pull the latest image
docker pull rodrigodiez/zorrohttp:latest

# Run zorro http server with memory storage
docker run -p 8080:8080 rodrigodiez/zorrohttp:latest --port 8080 --storage-driver memory

# Run zorro http server with BoltDB storage (initialises a new db if $BOLTDB_PATH does not exist)
docker run -p 8080:8080 rodrigodiez/zorrohttp:latest --port 8080 --storage-driver boltdb -storage-path $BOLTDB_PATH

# Run zorro http server with DynamoDB storage (requires tables to configured with the following key {ID: String})
docker run -p 8080:8080    -e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY rodrigodiez/zorrohttp:latest --port 8080 --storage-driver dynamodb -dynamodb-keys-table $DINAMODB_KEYS_TABLE -dynamodb-values-table $DINAMODB_VALUES_TABLE -aws-region $AWS_REGION

# Mask
curl -X POST http://localhost:8080/mask/<key>

# Unmask
curl -X POST http://localhost:8080/unmask/<value>
```

## Using Zorro as a package in your app
```go
package main

import (
	"fmt"

	"github.com/rodrigodiez/zorro/pkg/generator/uuid"
	"github.com/rodrigodiez/zorro/pkg/service"
	"github.com/rodrigodiez/zorro/pkg/storage/memory"
)

func main() {
	z := service.New(
		uuid.NewV4(),
		memory.New(),
	)

	value := z.Mask("foo")
	fmt.Println(value)
	// Will print something like '870284f9-c269-4175-8ab9-8e0a094a64ab'

	key, _ := z.Unmask(value)
	fmt.Println(key)
	// Will print 'foo'

	// Once generated masks are idempotent!
	value = z.Mask("foo")
	fmt.Println(value)
	// Will print same mask as before
}
```

## Documentation
- [Godoc](https://godoc.org/github.com/rodrigodiez/zorro) documentation is available.

## Operations
- Mask
- Unmask
- BatchMask (to-do)
- BatchUnmask (to-do)

## Servers
- HTTP (available)
- HTTPS (to-do)
- [GRPC](https://grpc.io/) (to-do)
- [Twirp](https://github.com/twitchtv/twirp) (to-do)

## Generators
- UUIDv4

## Storage
- In-Memory (available)
- [Bolt](https://github.com/boltdb/bolt) (available)
- [DynamoDB](https://aws.amazon.com/dynamodb/) (available)
- [Redis](https://redis.io/) (to-do)
- [MySQL](https://www.mysql.com/) (to-do)
- Chain (multiple storages) (to-do)

# Metrics
Zorro can emit some internal metrics

```go
package main

import (
	"expvar"

	"github.com/rodrigodiez/zorro/pkg/generator/uuid"
	"github.com/rodrigodiez/zorro/pkg/service"
	"github.com/rodrigodiez/zorro/pkg/storage"
	"github.com/rodrigodiez/zorro/pkg/storage/memory"
)

func main() {
	service.New(
		uuid.NewV4(),
		memory.New().WithMetrics(&storage.Metrics{
			LoadOps:    expvar.NewInt("loadOps"),
			StoreOps:   expvar.NewInt("storeOps"),
			ResolveOps: expvar.NewInt("resolveOps"),
		}),
	).WithMetrics(&service.Metrics{
		MaskOps:   expvar.NewInt("maskOps"),
		UnmaskOps: expvar.NewInt("unmaskOps"),
	})
}
```

## Contributing
If you want to contribute to the development of Zorro you are more than welcome!

- Fixes: Go ahead and create a PR! :D
- Enhancements: Have a look to the [open issues](https://github.com/rodrigodiez/zorro/issues). If your enhancement does not fit any of the existing ones please create a new issue and describe your use case so we can discuss how to make it real! :D

## About the author
My name is Rodrigo Díez Villamuera. Above anything else I am a passionate maker

I started Zorro as a way to increase my experience with Golang. High performance servers, cloud storage, concurrency... you name it!

I am also available to hire as a contractor. I am specialised in

- PHP/Go/Node development
- AWS Solutions
- Team leadership
- Agile development

If you need a hand or two... get in touch!
[Linkedin](https://www.linkedin.com/in/rodrigodiezvillamuera/) | [rodrigodiez.io](http://rodrigodiez.io)

## License
Zorro is free software and it is distributed under the terms and conditions of the [MIT License](https://choosealicense.com/licenses/mit/).
