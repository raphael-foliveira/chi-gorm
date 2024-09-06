package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type Cfg struct {
	DatabaseURL string
	JwtSecret   string
}

var configInstance *Cfg

func LoadCfg(path string) *Cfg {
	content, err := getFileContent(path)
	if err != nil {
		panic(err)
	}
	parseEnv(content)
	setEnvs()
	return configInstance
}

func Config() *Cfg {
	return &Cfg{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JwtSecret:   os.Getenv("JWT_SECRET"),
	}
}

func setEnvs() {
	configInstance = &Cfg{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JwtSecret:   os.Getenv("JWT_SECRET"),
	}
}

func parseEnv(s string) error {
	contentLines := strings.Split(s, "\n")
	for _, line := range contentLines {
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		pair := strings.SplitN(line, "=", 2)
		if len(pair) != 2 {
			log.Println(fmt.Sprintf("pair: %s", pair))
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
