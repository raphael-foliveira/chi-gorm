package internal

type Config struct {
	DatabaseURL string
}

var TestConfig = Config{
	DatabaseURL: "postgres://postgres:postgres@localhost:5432/test?sslmode=disable",
}
