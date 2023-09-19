package main

import (
	"fmt"

	"github.com/raphael-foliveira/chi-gorm/pkg/http/server"
	"github.com/raphael-foliveira/chi-gorm/pkg/persistence/sqlstore"
)

func main() {
	sqlstore.InitPg()
	err := server.Start()
	if err != nil {
		fmt.Println(err)
	}
}
