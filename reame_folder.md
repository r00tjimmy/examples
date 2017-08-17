# Micro Examples  [![Travis CI](https://travis-ci.org/micro/examples.svg?branch=master)](https://travis-ci.org/micro/examples) [![Go Report Card](https://goreportcard.com/badge/micro/examples)](https://goreportcard.com/report/github.com/micro/examples)

This is a repository for micro examples. Feel free to contribute.

## Contents

- [greeter](greeter) - A complete greeter example (includes python, ruby examples)
- [sidecar](sidecar) - Greeter service using the sidecar with multiple languages
- [booking](booking) - A booking.com demo application
- [plugins](plugins) - How to use plugins
- [broker](broker) - A example of using Broker for Publish and Subscribing.
- [client](client) - Usage of the Client package to call a service.
- [graceful](graceful) - Demonstrates graceful shutdown of a service
- [redirect](redirect) - An example of how to http redirect using an API service
- [server](server) - Use of the Server package directly to server requests.
- [service](service) - Example of the top level Service in go-micro.
- [shutdown](shutdown) - Demonstrates graceful shutdown via context cancellation
- [template](template) - Api, web and srv service templates generated with the 'micro new' command
- [grpc](grpc) - Examples of how to use [go-grpc](https://github.com/micro/go-grpc)
- [flags](flags) - Using command line flags with a service
- [command](command) - An example of bot commands as micro services
- [stream](stream) - An example of a streaming service and client

## External

- [auth-srv](https://github.com/micro/auth-srv) - An Oauth2 authentication service
- [geo-srv](https://github.com/micro/geo-srv) - Geolocation tracking service using hailocab/go-geoindex
- [geo-web](https://github.com/micro/geo-web) - Web demo for the geo srv
- [geo-api](https://github.com/micro/geo-api) - API for the geo srv
- [discovery-srv](https://github.com/micro/discovery-srv) - A discovery in the micro platform
- [geocode-srv](https://github.com/micro/geocode-srv) - A geocoding service using the Google Geocoding API
- [hailo-srv](https://github.com/micro/hailo-srv) - A service for the hailo taxi service developer api
- [monitor-srv](https://github.com/micro/monitor-srv) - A monitoring service for Micro services
- [place-srv](https://github.com/micro/place-srv) - A microservice to store and retrieve places (includes Google Place Search API)
- [slack-srv](https://github.com/micro/slack-srv) - The slack bot API as a go-micro RPC service
- [trace-srv](https://github.com/micro/trace-srv) - A distributed tracing microservice in the realm of dapper, zipkin, etc
- [twitter-srv](https://github.com/micro/twitter-srv) - A microservice for the twitter API
- [user-srv](https://github.com/micro/user-srv)	- A microservice for user management and authentication

## Dependencies

### Service Discovery

All services require service discovery. The default is Consul or MDNS.

#### Consul

Install
```
brew install consul
```

Run
```
consul agent -dev
```

#### Multicast DNS

Use flag `--registry=mdns` for a zero dependency configuration

### Protobuf

Protobuf is used for code generation of message types and client/hander stubs.

If making changes recompile the protos.

#### Install
```shell
go get github.com/micro/protobuf/{proto,protoc-gen-go}
```

#### Compile Proto

```shell
protoc -I$GOPATH/src --go_out=plugins=micro:$GOPATH/src \
        $GOPATH/src/github.com/micro/examples/service/proto/greeter.proto
```


