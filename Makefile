GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=out/resumator-service

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME)
test:
	$(GOTEST) -v ./...
run: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)


