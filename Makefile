setup:
	go mod tidy

build:
	go build -o bin/gendiff ./cmd/main.go
	
lint:
	golangci-lint run ./code/...
