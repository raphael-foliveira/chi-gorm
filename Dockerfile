FROM golang:alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/app

CMD ["/app/main"]
