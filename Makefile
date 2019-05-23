# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=bin/gopress
CMD_DIR=cmd/gopress/*
INSTALLBINDIR := /usr/local/bin
 
.PHONY: install

all: clean build install

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v $(CMD_DIR)
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v $(CMD_DIR)
	./$(BINARY_NAME)

install:
	cp ./$(BINARY_NAME) $(INSTALLBINDIR)
