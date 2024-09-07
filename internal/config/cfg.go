package config

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

var (
	DatabaseURL string
	JwtSecret   string
)

func Initialize(path ...string) {
	if len(path) > 0 {
		content, err := getFileContent(path[0])
		if err != nil {
			panic(err)
		}
		parseEnv(content)
	}
	Load()
}

func Load() {
	DatabaseURL = os.Getenv("DATABASE_URL")
	JwtSecret = os.Getenv("JWT_SECRET")
}

func parseEnv(s string) error {
	contentLines := strings.Split(s, "\n")
	for _, line := range contentLines {
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		pair := strings.SplitN(line, "=", 2)
		if len(pair) != 2 {
			slog.Error(fmt.Sprintf("pair: %s", pair))
			return ErrMalformedEnvEntry
		}
		key := pair[0]
		val := strings.Join(pair[1:], "=")
		os.Setenv(key, val)
	}
	return nil
}

var ErrMalformedEnvEntry = errors.New("malformed env entry")

func getFileContent(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	content := string(bytes)
	return removeQuotes(content), nil
}

func removeQuotes(content string) string {
	return strings.ReplaceAll(content, `"`, "")
}
