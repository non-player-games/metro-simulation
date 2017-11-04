# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_FOLDER=bin
BINARY_NAME=simulation
BINARY_UNIX=$(BINARY_NAME)_unix
MAIN_FILE=cmd/cli/main.go

all: install test build
install:
	$(GOGET) -t ./...
build: 
	$(GOBUILD) -o $(BINARY_FOLDER)/$(BINARY_NAME) -v $(MAIN_FILE)
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_FOLDER)/$(BINARY_NAME)
	rm -f $(BINARY_FOLDER)/$(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_FOLDER)/$(BINARY_NAME) -v $(MAIN_FILE)
	./$(BINARY_FOLDER)/$(BINARY_NAME)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_FOLDER)/$(BINARY_UNIX) -v $(MAIN_FILE)
