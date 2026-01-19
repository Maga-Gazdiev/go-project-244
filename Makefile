setup:
	cd code && go mod tidy

build:
	cd code && go build -o ../bin/gendiff ./cmd/gendiff
	
lint:
	cd code && golangci-lint run ./...