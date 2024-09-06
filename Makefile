run:
	go run cmd/main.go
	
cover:
	go tool cover -html=c.out;

test:
	docker compose up -d database && \
	go test ./... -cover -coverpkg=../... -coverprofile=c.out;
	docker compose stop;

test-cover: test cover

test-watch:
	air -c .air.test.toml

build:
	GOOS=linux go build -o bin/main cmd/main.go
	GOOS=windows go build -o bin/main.exe cmd/main.go

dev:
	air

docker-up:
	docker compose up -d --build

docker-down:
	docker compose down
