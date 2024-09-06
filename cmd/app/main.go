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
		Addr:    ":3000",
		Handler: mux,
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)

	log.Println("server starting on port 3000")
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	<-ch

	log.Println("server interrupted")
}
