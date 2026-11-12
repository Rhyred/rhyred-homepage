package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Rhyred/rhyred-homepage/internal/config"
	"github.com/Rhyred/rhyred-homepage/internal/handlers"
	"github.com/Rhyred/rhyred-homepage/internal/router"
	"github.com/Rhyred/rhyred-homepage/internal/services"
)

func main() {
	cfg := config.Load()

	// Initialize data service
	dataSvc := services.NewDataService(cfg.DataDir)
	if err := dataSvc.Load(); err != nil {
		log.Fatalf("Gagal to load data: %v", err)
	}

	// Initialize handlers
	h := handlers.NewHandler(dataSvc)

	// Setup router
	mux := router.New(h, "web/static")

	// Server with timeouts
	addr := ":" + cfg.Port
	srv := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Server banner
	fmt.Fprintf(os.Stdout, "\n  ╔══════════════════════════════════════╗\n")
	fmt.Fprintf(os.Stdout, "  ║  RHYRED.NET — Development Server     ║\n")
	fmt.Fprintf(os.Stdout, "  ║  http://localhost%-20s║\n", addr)
	fmt.Fprintf(os.Stdout, "  ╚══════════════════════════════════════╝\n\n")

	// Graceful shutdown channel
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Server starting on %s", addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	<-stop
	log.Println("Maehan server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server Modar paksa: %v", err)
	} else {
		log.Println("Server ereun normal.")
	}
}
