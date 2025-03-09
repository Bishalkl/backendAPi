build:
	@go build -o bin/backendAPi cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/backendAPi