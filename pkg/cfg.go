package pkg

import "os"

type Config struct {
	DatabaseURL string
}

var MainConfig = Config{
	DatabaseURL: os.Getenv("DATABASE_URL"),
}

var TestConfig = Config{
	DatabaseURL: "postgres://postgres:postgres@localhost:5432/test?sslmode=disable",
}
