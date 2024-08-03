build:
	@go build -o bin/mahder

run: build
	@./bin/mahder

test:
	@go test ./... -v