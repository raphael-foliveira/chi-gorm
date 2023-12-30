package cfg

import (
	"os"
	"strings"
)

type Config struct {
	DatabaseURL string
}

var MainConfig Config
var TestConfig Config

func LoadCfg(envFile string) error {
	bytes, err := os.ReadFile(envFile)
	if err != nil {
		return err
	}
	content := string(bytes)
	content = strings.ReplaceAll(content, `"`, "")
	contentLines := strings.Split(content, "\n")
	for _, line := range contentLines {
		pair := strings.Split(line, "=")
		os.Setenv(pair[0], pair[1])
	}
	MainConfig = Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	TestConfig = Config{
		DatabaseURL: os.Getenv("TEST_DATABASE_URL"),
	}
	return nil
}
