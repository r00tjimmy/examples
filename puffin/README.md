# Puffin

An data packing application, get data from API, database, etc.

## Contents

- **srv** - an RPC puffin service
- **api** -  RPC API and RESTful API
- **web** - how to use go-web to write web services

## Deps

Service discovery is required for all services. Default is Consul or MDNS. You can also use plugins from 
[micro/plugins](https://github.com/micro/go-plugins).

### MDNS

Use the flag `--registry=mdns`

### Consul

```
brew install consul
consul agent -dev
```

## Run Service

Start go.micro.srv.puffin
```shell
go run srv/main.go
```


## API

HTTP based requests can be made via the micro API. Micro logically separates API services from backend services. By default the micro API 
accepts HTTP requests and converts to *api.Request and *api.Response types. Find them here [micro/api/proto](https://github.com/micro/micro/tree/master/api/proto).

Run the go.micro.api.puffin API Service
```shell
go run api/api.go 
```

Run the micro API
```shell
micro api
```

Call go.micro.api.puffin via API
```go
curl http://localhost:8080/puffin/say/hello?name=John
```

Examples of other API handlers can be found in the API directory.

## Sidecar

The sidecar is a language agnostic RPC proxy.

Run the micro sidecar
```shell
micro sidecar
```

Call go.micro.srv.puffin via sidecar
```shell
curl -H 'Content-Type: application/json' -d '{"name": "john"}' http://localhost:8081/puffin/say/hello
```

The sidecar provides all the features of go-micro as a HTTP API. Learn more at [micro/car](https://github.com/micro/micro/tree/master/car).

## Web

The micro web is a web dashboard and reverse proxy to run web apps as microservices.

Run go.micro.web.puffin
```
go run web/web.go 
```

Run the micro web
```shell
micro web
```

Browse to http://localhost:8082/puffin
