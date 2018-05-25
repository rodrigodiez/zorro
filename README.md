# Zorro

[![build](	https://img.shields.io/travis/rodrigodiez/zorro.svg)](https://travis-ci.org/rodrigodiez/zorro)
[![Go Report Card](https://goreportcard.com/badge/github.com/rodrigodiez/zorro)](https://goreportcard.com/report/github.com/rodrigodiez/zorro)
[![](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/rodrigodiez/zorro)
[![](https://images.microbadger.com/badges/image/rodrigodiez/zorro-http.svg)](https://microbadger.com/images/rodrigodiez/zorro-http "Get your own image badge on microbadger.com")
[![MIT License](https://img.shields.io/github/license/rodrigodiez/zorro.svg)](https://github.com/rodrigodiez/zorro/blob/master/LICENSE.md)

Zorro is a service that allows masking and unmasking of text over multiple transports, storage drivers and masking strategies.

Zorro comes with some reference servers but can also be used as a package in your own application.

[![gopher](https://github.com/egonelbre/gophers/raw/master/.thumb/vector/superhero/standing.png)](https://github.com/egonelbre/gophers)

by [@egonelbre](https://github.com/egonelbre/gophers)

---

> **Important**: Zorro is under heavy development at the moment and its usage in production is **not** recommended

## Use cases
- Services that want to protect their private IDs

## Zorro in action
The easiest way to test Zorro is by running a Docker container of `zorro-http`, Zorro's http server.

```bash
# Pull the latest image
docker pull rodrigodiez/zorro-http:latest

# Run zorro http server with memory storage
docker run -p 8080:8080 rodrigodiez/zorro-http:latest --port 8080 --storage-driver memory --debug

# Run zorro http server with BoltDB storage (initialises a new db if $BOLTDB_PATH does not exist)
docker run -p 8080:8080 rodrigodiez/zorro-http:latest --port 8080 --storage-driver boltdb -storage-path $BOLTDB_PATH --debug

# Run zorro http server with DynamoDB storage (requires tables to configured with the following key {ID: String})
docker run -p 8080:8080    -e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY rodrigodiez/zorro-http:latest --port 8080 --storage-driver dynamodb -dynamodb-keys-table $DINAMODB_KEYS_TABLE -dynamodb-values-table $DINAMODB_VALUES_TABLE -aws-region $AWS_REGION --debug

# Mask
curl -X POST http://localhost:8080/mask/<key>

# Unmask
curl -X POST http://localhost:8080/unmask/<value>

# Metrics
curl http://localhost:8080/debug/vars
```

## Installation
```
go get -u github.com/rodrigodiez/zorro/...
```

## Operations
- Mask
- Unmask
- BatchMask
- BatchUnmask

## Servers
- HTTP (available `zorro-http`)
- [GRPC](https://grpc.io/) (available `zorro-grpc`, not tls support yet)
- HTTPS (to-do)
- [Twirp](https://github.com/twitchtv/twirp) (to-do)

## Masking strategies
- UUIDv4

## Storage
- In-Memory (available)
- Bolt (available)
- DynamoDB (available)
- Redis (to-do)
- MySQL (to-do)

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
