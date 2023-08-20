test:
	go test ./... -coverprofile=coverage.out;
	go tool cover -html=coverage.out -html=c.out;

run:
	go run cmd/main.go

build:
	go build -o bin/main cmd/main.go