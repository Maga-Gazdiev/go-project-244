build:
	cd code && go build -o ../bin/gendiff ./cmd/gendiff
	
lint:
	cd code && golangci-lint run ./...