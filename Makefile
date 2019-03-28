default: build

fmt:
	./build-tools/scripts/check --fmt

build:
	./build-tools/scripts/build

run_server:
	./bin/simple-client-server_darwin_amd64 --loglvl debug --port 9000

run_client:
	./bin/simple-client-server_darwin_amd64 client --msg "digital ocean rocks"

check_fmt:
	./build-tools/scripts/check

release:
	docker build -f build-tools/res/Dockerfile -t jharshman/simple-client-server ./bin
	# docker login -u $(DOCKER_USER) -p $(DOCKER_PASSWORD)

install:
	./build-tools/scripts/install -c $(CTX) -n $(NS)

clean:
	rm -f ./bin/*
