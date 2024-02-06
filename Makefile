build:
	@go build -o bin/api

run:
	@./bin/api

test:
	@got test -v ./...