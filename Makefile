default: build

fmt:
	./build-tools/scripts/check --fmt

build:
	./build-tools/scripts/build

run_server:
	./bin/simple-client-server_darwin_amd64 --loglvl debug --port 8081

run_client:
	./bin/simple-client-server_darwin_amd64 client

check_fmt:
	./build-tools/scripts/check

release:
	docker build -t jharshman/simple-client-server .
	# docker login -u $(DOCKER_USER) -p $(DOCKER_PASSWORD)

install:
	./build-tools/scripts/install -c $(CTX) -n $(NS)

clean:
	rm -f ./bin/*
