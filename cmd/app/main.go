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
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

func main() {
	if err := database.Start(); err != nil {
		log.Fatal("database.Start failed:", err)
	}
	defer database.Close()

	mux := chi.NewMux()

	clientsRepo := repository.NewClients(database.DB)
	ordersRepo := repository.NewOrders(database.DB)
	productsRepo := repository.NewProducts(database.DB)

	clientsController := controller.NewClients(clientsRepo, ordersRepo)
	productsController := controller.NewProducts(productsRepo)
	ordersController := controller.NewOrders(ordersRepo)
	healthCheckController := controller.NewHealthCheck()

	mux.Mount("/clients", clientsController.Routes())
	mux.Mount("/products", productsController.Routes())
	mux.Mount("/orders", ordersController.Routes())
	mux.Mount("/health-check", healthCheckController.Routes())

	s := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("server starting on port 3000")
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-ch

	log.Println("server interrupted")
}
