package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
		log.Fatalf("Failed to load data: %v", err)
	}

	// Initialize handlers
	h := handlers.NewHandler(dataSvc)

	// Setup router
	mux := router.New(h, "web/static")

	// Start server
	addr := ":" + cfg.Port
	fmt.Fprintf(os.Stdout, "\n  ╔══════════════════════════════════════╗\n")
	fmt.Fprintf(os.Stdout, "  ║  RHYRED.NET — Development Server     ║\n")
	fmt.Fprintf(os.Stdout, "  ║  http://localhost%-20s║\n", addr)
	fmt.Fprintf(os.Stdout, "  ╚══════════════════════════════════════╝\n\n")

	log.Printf("Server starting on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

