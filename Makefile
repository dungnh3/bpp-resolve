BIN_FOLDER = bin
SERVER = runtime
GO_CMD_RUNTIME = cmd/main.go

build:
	GOSUMDB=off go mod tidy
	go build -o $(BIN_FOLDER)/$(SERVER) $(GO_CMD_RUNTIME)

lint:
	golangci-lint run --allow-parallel-runners

mock:
	go get -u github.com/stretchr/testify/mock
	go get -u github.com/vektra/mockery/.../
	cd internal && mockery -all -inpkg -keeptree --case underscore && cd ..

test:
	go test -race -count=1 -v -cover ./...

start:
	chmod +x start.sh
	./start.sh


# BUILD DOCKER IMAGE
IMAGE_NAME_SERVICE = bpp
VERSION_SERVICE = 0.0.1

build-docker-image:
	docker build -t $(IMAGE_NAME_SERVICE):$(VERSION_SERVICE) --force-rm -f Dockerfile .

run-docker-image:
	docker run -it --name $(IMAGE_NAME_SERVICE) -p 8080:8080 $(IMAGE_NAME_SERVICE):$(VERSION_SERVICE)

.PHONY: build lint mock test build-docker-image run-docker-image