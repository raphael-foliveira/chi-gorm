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

func init() {
	Load(".env")
}

func Load(path ...string) {
	if len(path) > 0 {
		content, err := getFileContent(path[0])
		if err != nil {
			panic(fmt.Errorf("error loading env file (path: %s): %w", path[0], err))
		}
		if err := parseEnv(content); err != nil {
			panic(err)
		}
	}
	load()
}

func load() {
	DatabaseURL = getEnv("DATABASE_URL", "")
	JwtSecret = getEnv("JWT_SECRET", "super secret")
}

func getEnv(key, dflt string) string {
	val := os.Getenv(key)
	if val == "" {
		return dflt
	}
	return val
}

func parseEnv(s string) error {
	contentLines := strings.Split(s, "\n")
	for _, line := range contentLines {
		if isEmptyOrComment(line) {
			continue
		}
		pair := strings.SplitN(line, "=", 2)
		if len(pair) != 2 {
			slog.Error(fmt.Sprintf("pair: %s", pair))
			return errors.New("malformed env entry")
		}
		key := pair[0]
		val := strings.Join(pair[1:], "=")
		os.Setenv(key, val)
	}
	return nil
}

func isEmptyOrComment(line string) bool {
	return line == "" || strings.HasPrefix(line, "#")
}

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
