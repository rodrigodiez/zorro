# Zorro
Zorro is a complete microservice and a set of libraries to mask and unmask IDs focused on speed, reliability and flexibility.

> **Important**: Zorro is under heavy development at the moment and under any circumstance its usage in production is recommended.

# Use cases
- Services that expose objects which internal IDs must be kept private but resolvable.

# Run the demo
> `zorrohttp` will be installed in your `$GOPATH/bin`
```
go get -u github.com/rodrigodiez/zorro/cmd/zorrohttp
zorrohttp --port 8080

curl http://localhost:8080/mask/foo
> 695bcafd-2fa4-4743-b30a-b42cf853fcd3

curl http://localhost:8080/unmask/695bcafd-2fa4-4743-b30a-b42cf853fcd3
> foo
```

# Documentation
- [Godoc](https://godoc.org/golang.org/github.com/rodrigodiez/zorro)

# Servers
> Developers can create their own servers

- HTTP (in progress, demo available)
- HTTPS (to-do)
- GRPC (to-do)
- Twirp (to-do)

# Generators
> Developers can create their own generators
- UUIDv4

# Storage
> Developers can create their own storages
- In-Memory
- Bolt (to-do)
- DynamoDB (to-do)
- Redis (to-do)
- MySQL (to-do)