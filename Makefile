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
	docker build -f build-tools/res/Dockerfile -t jharshman/simple-client-server ./bin

publish: build release
	docker login -u $(DOCKER_USER) -p $(DOCKER_PASSWORD)
	docker push jharshman/simple-client-server:latest

protobuf:
	protoc --proto_path=pkg/grpc --go_out=plugins=grpc:pkg/grpc echo.proto

mockgen:
	mockgen --source=pkg/grpc/echo.pb.go --destination=mock/echo.go --package=mock

clean:
	rm -rf ./bin/*
