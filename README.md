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
