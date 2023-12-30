package main

import (
	"fmt"

	"github.com/raphael-foliveira/chi-gorm/internal/cfg"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/http/server"
)

func main() {
	err := cfg.LoadEnv(".env")
	if err != nil {
		panic(err)
	}
	fmt.Println("configs: ", cfg.MainConfig, cfg.TestConfig)
	err = database.InitDb(cfg.MainConfig.DatabaseURL)
	if err != nil {
		panic(err)
	}
	err = server.NewServer().Start()
	if err != nil {
		panic(err)
	}
}
