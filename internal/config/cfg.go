package config

import (
	"os"
	"strings"
)

type cfg struct {
	DatabaseURL string
	JwtSecret   string
}

var configInstance *cfg

func LoadCfg(path string) {
	content, err := getFileContent(path)
	if err != nil {
		panic(err)
	}
	parseEnv(content)
	setEnvs()
}

func Config() *cfg {
	return &cfg{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JwtSecret:   os.Getenv("JWT_SECRET"),
	}
}

func setEnvs() {
	configInstance = &cfg{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JwtSecret:   os.Getenv("JWT_SECRET"),
	}
}

func parseEnv(s string) {
	contentLines := strings.Split(s, "\n")
	for _, line := range contentLines {
		pair := strings.Split(line, "=")
		if len(pair) > 1 {
			key := pair[0]
			val := strings.Join(pair[1:], "=")
			os.Setenv(key, val)
		}
	}
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
