NAME     := dialog-stress-bots
REVISION := $(shell git rev-parse --short HEAD)

GOPATH=$(shell go env GOPATH)

default: deps build

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf vendor/*

.PHONY: dep
dep:
ifeq ($(shell command -v dep 2> /dev/null),)
	go get -u -v github.com/golang/dep/cmd/dep
endif

.PHONY: deps
deps: dep
	dep ensure -v

.PHONY: build
build:
	go fmt .
	go build -o bin/$(NAME) ./main.go

.PHONY: docker
docker:
	docker build . -t $(NAME)
