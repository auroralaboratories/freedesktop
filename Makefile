all: deps fmt test

deps:
	@go list golang.org/x/tools/cmd/goimports || go get golang.org/x/tools/cmd/goimports
	go get ./...

fmt:
	goimports -w .
	go vet .

test:
	go test -race ./...

