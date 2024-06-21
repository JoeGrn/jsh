GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
GOFMT = $(GOCMD) fmt

DOCKERCMD = docker
DOCKERBUILD = $(DOCKERCMD) build
DOCKERRUN = $(DOCKERCMD) run

BINARY_NAME = jsh
BUILD_DIR = dist
ENTRY_POINT = ./cmd/jsh

all: test build

build:
	$(GOBUILD) -o ./${BUILD_DIR}/$(BINARY_NAME) -v $(ENTRY_POINT)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f  -r ./$(BUILD_DIR)

run:
	$(GOBUILD) -o ./${BUILD_DIR}/$(BINARY_NAME) -v $(ENTRY_POINT)
	./${BUILD_DIR}/$(BINARY_NAME)

fmt:
	$(GOFMT) ./...

docker-build:
	${DOCKERBUILD} -t jsh .

docker-run:
	${DOCKERRUN} -p 8080:8080 jsh
