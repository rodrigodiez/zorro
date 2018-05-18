# Zorro

![build](https://travis-ci.org/rodrigodiez/zorro.svg?branch=master)

Zorro is a microservice to mask and unmask IDs focused on speed, reliability and flexibility. 

![gopher](https://github.com/egonelbre/gophers/raw/master/.thumb/vector/superhero/standing.png)

Gopher by [@egonelbre](https://github.com/egonelbre/gophers)

---

> **Important**: Zorro is under heavy development at the moment and under any circumstance its usage in production is recommended

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

## License
Zorro is free software license under the [MIT License](https://choosealicense.com/licenses/mit/) terms and conditions.
