cover:
	go tool cover -html=c.out;

test:
	go test ./... -v -cover -coverpkg=../... -coverprofile=c.out;

run:
	go run cmd/main.go

build:
	GOOS=linux go build -o bin/main cmd/main.go
	GOOS=windows go build -o bin/main.exe cmd/main.go

dev:
	air