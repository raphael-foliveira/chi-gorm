cover:
	go tool cover -html=c.out;

test:
	go test ./... -coverprofile=c.out;

run:
	go run cmd/main.go

build:
	go build -o bin/main cmd/main.go

dev:
	air