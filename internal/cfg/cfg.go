package cfg

import (
	"os"
	"strings"
)

type cfg struct {
	DatabaseURL string
}

var Cfg cfg

func LoadCfg(path string) error {
	content, err := getFileContent(path)
	if err != nil {
		return err
	}
	parseEnv(content)
	Cfg.DatabaseURL = os.Getenv("DATABASE_URL")
	return nil
}

func parseEnv(s string) {
	contentLines := strings.Split(s, "\n")
	for _, line := range contentLines {
		pair := strings.Split(line, "=")
		if len(pair) > 1 {
			os.Setenv(pair[0], pair[1])
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
