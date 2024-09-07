package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
)

func main() {
	if err := database.Start(); err != nil {
		log.Fatal("database.Start failed:", err)
	}
	defer database.Close()

	mux := chi.NewMux()

	controller.Mount(mux)

	s := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	log.Println("server starting on port 3000")
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-ch

	log.Println("server interrupted")
}
