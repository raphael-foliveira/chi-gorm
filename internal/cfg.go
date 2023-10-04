package pkg

import "os"

type Config struct {
	DatabaseURL string
}

var MainConfig = Config{
	DatabaseURL: os.Getenv("DATABASE_URL"),
}

var TestConfig = Config{
	DatabaseURL: os.Getenv("TEST_DATABASE_URL"),
}
