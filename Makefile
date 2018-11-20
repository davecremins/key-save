.PHONY: all test coverage format

all: get build install

get:
	go get ./...

build:
	go build ./...

install:
	go install ./...

format:
	gofmt -w .

test:
	go test ./... -v -coverprofile .coverage.txt
	go tool cover -func .coverage.txt

coverage: test
	go tool cover -html=.coverage.txt
