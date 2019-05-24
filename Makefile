# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=bin/gopress
CMD_DIR=cmd/gopress/*
INSTALLBINDIR := /usr/local/bin
 
.PHONY: install

all: clean test build install

test:
	$(GOTEST) -v ./...

build: 
	$(GOBUILD) -o $(BINARY_NAME) $(CMD_DIR)

clean: 
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) $(CMD_DIR)
	./$(BINARY_NAME)

install:
	cp ./$(BINARY_NAME) $(INSTALLBINDIR)
