# Go parameters
GOCMD=go 
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=heroes
BINARY_UNIX=$(BINARY_NAME)_unix 

all: test build
build: 
	$(GOBUILD) -o $(BINARY_NAME)
test:
	$(GOTEST)
run:
	$(GOBUILD) -o $(BINARY_NAME) 
	./$(BINARY_NAME)