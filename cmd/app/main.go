package main

import (
	"fmt"

	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/http/server"
)

func main() {
	database.InitPg()
	err := server.Start()
	if err != nil {
		fmt.Println(err)
	}
}
