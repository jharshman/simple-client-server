default: build

fmt:
	./build-tools/scripts/check --fmt

build:
	./build-tools/scripts/build

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
	# docker login -u $(DOCKER_USER) -p $(DOCKER_PASSWORD)

install:
	./build-tools/scripts/install -c $(CTX) -n $(NS)

clean:
	rm -rf ./bin/*
