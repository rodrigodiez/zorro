# Zorro

![build](https://travis-ci.org/rodrigodiez/zorro.svg?branch=master)

Zorro allows developers to perform two way key/value lookups.

Zorro can be used with one of the provided servers or as a golang package for maximum flexibility.

![gopher](https://github.com/egonelbre/gophers/raw/master/.thumb/vector/superhero/standing.png)

Gopher by [@egonelbre](https://github.com/egonelbre/gophers)

---

> **Important**: Zorro is under heavy development at the moment and its usage in production is **not** recommended

## Use cases
- Services that want to protect their private IDs by translating them to public ones while keeping the ability to translate them back

## Using Zorro as a server

> At the moment only an http server is available

```bash
# Install zorrohttp
go get -u github.com/rodrigodiez/zorro/cmd/zorrohttp

# Choose your port and storage!
zorrohttp --port 8080 --storage-driver memory
zorrohttp --port 8080 --storage-driver boltdb -storage-path /tmp/elzorro.db

# Mask an key
curl -X POST http://localhost:8080/mask/<key>

# Unmask a key
curl -X POST http://localhost:8080/unmask/<value>
```

## Using Zorro as a package
```go
package main

import (
	"fmt"

	"github.com/rodrigodiez/zorro"
)

func main() {
	z := zorro.New(
		zorro.NewUUIDv4Generator(),
		zorro.NewInMemoryStorage(),
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
> Developers can create their own servers

- HTTP (available)
- HTTPS (to-do)
- [GRPC](https://grpc.io/) (to-do)
- [Twirp](https://github.com/twitchtv/twirp) (to-do)

## Generators
> Developers can create their own generators
- UUIDv4

## Storage
> Developers can create their own storages
- In-Memory (available)
- [Bolt](https://github.com/boltdb/bolt) (available)
- [DynamoDB](https://aws.amazon.com/dynamodb/) (to-do)
- [Redis](https://redis.io/) (to-do)
- [MySQL](https://www.mysql.com/) (to-do)
- Chain (multiple storages) (to-do)

## Contributing
If you want to contribute to the development of Zorro you are more than welcome!

- Fixes: Go ahead and create a PR! :D
- Enhancements: Have a look to the [open issues](https://github.com/rodrigodiez/zorro/issues). If your enhancement does not fit any of the existing ones please create a new issue and describe your use case so we can discuss how to make it real! :D

## License
Zorro is free software and it is distributed under the terms and conditions of the [MIT License](https://choosealicense.com/licenses/mit/).
