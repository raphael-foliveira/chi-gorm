run:
	go run cmd/main.go
	
test-unit:
	go test ./... -tags=unit -v -cover -coverpkg=../... -coverprofile=c.out;

test-integration:
	docker compose up -d database && \
	go test ./... -tags=integration -cover -v -coverpkg=../... -coverprofile=c.out;
	docker compose stop;

cover:
	go tool cover -html=c.out;

test-all: 
	docker compose up -d database && \
	go test ./... -tags="integration,unit" -v -cover -coverpkg=../... -coverprofile=c.out;
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
