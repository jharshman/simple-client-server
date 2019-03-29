[![Build Status](https://travis-ci.org/jharshman/simple-client-server.svg?branch=master)](https://travis-ci.org/jharshman/simple-client-server)
[![Go Report Card](https://goreportcard.com/badge/github.com/jharshman/simple-client-server)](https://goreportcard.com/report/github.com/jharshman/simple-client-server)

# simple-client-server
simple client and server

## Run with Docker
You can use the Pre-built image, just remember to generate some certs first. Either specify your own or use `make gencerts`.
```
$ docker run -d --rm --name echo_server -v $PWD/bin/tls:/tls jharshman/simple-client-server:latest server --loglvl debug --cert /tls/cert.pem --key /tls/key.pem

$ docker run --rm -v $PWD/bin/tls:/tls --link=echo_server:echo.example.int jharshman/simple-client-server:latest client --client-cert /tls/cert.pem --server echo.example.int --msg "knock knock" 
```

## Build
```
$ make build
```

## Run the server
```
$ make run_server
```

## Run the client
```
$ make run_client
```

# Component Breakdown
## Server Usage
```
➜  simple-client-server git:(master) ./bin/simple-client-server_darwin_amd64 server --help
Run the echo server

Usage:
  simple-client-server server [flags]

Flags:
      --cert string     tls cert (default "cert.pem")
  -h, --help            help for server
      --key string      tls key (default "key.pem")
      --loglvl string   log level (default "info")
      --port string     grpc bind address (default "9000")
```

## Client Usage
```
➜  simple-client-server git:(master) ✗ ./bin/simple-client-server_darwin_amd64 client --help

Facilitates communication to Echo Server.

Usage:
  simple-client-server client [flags]

Flags:
      --client-cert string   client tls certificate (default "cert.pem")
  -h, --help                 help for client
      --msg string           message to send to echo server
      --server string        address of echo server (default "127.0.0.1")
      --server-port string   port of echo server (default "9000")
```