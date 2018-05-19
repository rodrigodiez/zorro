# Zorro

![build](https://travis-ci.org/rodrigodiez/zorro.svg?branch=master)

Zorro is a microservice to mask and unmask IDs focused on speed, reliability and flexibility. 

![gopher](https://github.com/egonelbre/gophers/raw/master/.thumb/vector/superhero/standing.png)

Gopher by [@egonelbre](https://github.com/egonelbre/gophers)

---

> **Important**: Zorro is under heavy development at the moment and its usage in production is **not** recommended

## Use cases
- Services that expose object which IDs must be kept private

## Run the demo
> A `zorrohttp` binary will be installed in your `$GOPATH/bin`
```
go get -u github.com/rodrigodiez/zorro/cmd/zorrohttp
zorrohttp --port 8080

curl http://localhost:8080/mask/foo
> 695bcafd-2fa4-4743-b30a-b42cf853fcd3

curl http://localhost:8080/unmask/695bcafd-2fa4-4743-b30a-b42cf853fcd3
> foo
```

## Documentation
- [Godoc](https://godoc.org/github.com/rodrigodiez/zorro)

## Operations
- Mask
- Unmask
- BatchMask (to-do)
- BatchUnmask (to-do)

## Servers
> Developers can create their own servers

- HTTP (in progress, demo available)
- HTTPS (to-do)
- [GRPC](https://grpc.io/) (to-do)
- [Twirp](https://github.com/twitchtv/twirp) (to-do)

## Generators
> Developers can create their own generators
- UUIDv4

## Storage
> Developers can create their own storages
- In-Memory
- [Bolt](https://github.com/boltdb/bolt) (to-do)
- [DynamoDB](https://aws.amazon.com/dynamodb/) (to-do)
- [Redis](https://redis.io/) (to-do)
- [MySQL](https://www.mysql.com/) (to-do)

## Contributing
If you want to contribute to the development of Zorro you are more than welcome!

- Fixes: Go ahead and create a PR! :D
- Enhancements: Have a look to the [open issues](https://github.com/rodrigodiez/zorro/issues). If your enhancement does not fit any of the existing ones please create a new issue and describe your use case so we can discuss how to make it real! :D

## License
Zorro is free software and it is distributed under the terms and conditions of the [MIT License](https://choosealicense.com/licenses/mit/).
