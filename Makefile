default: build

fmt:
	./build-tools/scripts/check --fmt

build:
	./build-tools/scripts/build -p 3

gencerts:
	./build-tools/scripts/gencerts

run_server: gencerts
	./bin/simple-client-server_darwin_amd64 server --loglvl debug --port 9000 --cert ./bin/tls/cert.pem --key ./bin/tls/key.pem

run_client:
	./bin/simple-client-server_darwin_amd64 client --client-cert ./bin/tls/cert.pem --msg "digital ocean rocks"

check_fmt:
	./build-tools/scripts/check

release:
	./build-tools/scripts/release $(DOCKER_USER) $(DOCKER_PASSWORD)

protobuf:
	protoc --proto_path=pkg/grpc --go_out=plugins=grpc:pkg/grpc echo.proto

mockgen:
	mockgen --source=pkg/grpc/echo.pb.go --destination=mock/echo.go --package=mock

clean:
	rm -rf ./bin/*
