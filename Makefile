build:
	@go build -o bin/api

run: build
	@./bin/api

test:
	@got test -v ./...