BUILD := $(shell git rev-parse --short HEAD)
NAME := $(shell basename "$(PWD)")
VERSION := $(shell git describe --tags)

BIN := $(GOPATH)/bin/$(NAME)
LDFLAGS=-ldflags "-s -w -X=main.build=$(BUILD) -X=main.name=$(NAME) -X=main.version=$(VERSION)"

all : compile run

benchmark :
	@ab -n 1000 -c 1000 http://localhost:3000/

compile :
	@go build $(LDFLAGS) -o $(BIN) main.go

pprof :
	@go tool pprof out.dump

profile :
	@curl -s http://localhost:3000/debug/pprof/profile?seconds=10 > out.dump

request :
	@curl -s http://localhost:3000/

run :
	$(BIN)
