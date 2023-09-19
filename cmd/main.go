package main

import (
	"fmt"

	"github.com/raphael-foliveira/chi-gorm/pkg/database"
	"github.com/raphael-foliveira/chi-gorm/pkg/http/server"
)

func main() {
	database.InitPg()
	err := server.Start()
	if err != nil {
		fmt.Println(err)
	}
}
