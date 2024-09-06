package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

func main() {
	config.Initialize()
	database.Initialize(config.DatabaseURL)
	defer database.Close()
	repository.Initialize()
	service.Initialize()
	controller.Initialize()

	mux := chi.NewMux()

	controller.Mount(mux)

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
