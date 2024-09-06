package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/container"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
)

func main() {
	cfg := config.Config()
	mux := container.InitializeDependencies(cfg)
	defer database.Close()

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	<-ch

	log.Println("server interrupted")
}
