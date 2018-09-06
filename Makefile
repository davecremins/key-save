format:
	gofmt -w .

test: format
	go test -v -cover ./...
