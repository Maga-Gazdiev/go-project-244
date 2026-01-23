setup:
	go mod tidy

build:
	go build -o bin/gendiff ./code/cmd/gendiff
	
lint:
	golangci-lint run ./code/...
