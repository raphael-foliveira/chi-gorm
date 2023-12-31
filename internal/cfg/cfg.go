package cfg

import (
	"os"
	"strings"
)

var DatabaseURL string

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
		if len(pair) > 1 {
			os.Setenv(pair[0], pair[1])
		}
	}
	DatabaseURL = os.Getenv("DATABASE_URL")
	return nil
}
