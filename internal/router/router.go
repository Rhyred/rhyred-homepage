package router

import (
	"net/http"

	"github.com/Rhyred/rhyred-homepage/internal/handlers"
)

// New creates and configures the HTTP router.
func New(h *handlers.Handler, staticDir string) http.Handler {
	mux := http.NewServeMux()

	// Pages
	mux.HandleFunc("GET /", h.Home)
	mux.HandleFunc("GET /about", h.About)
	mux.HandleFunc("GET /projects", h.Projects)
	mux.HandleFunc("GET /projects/{slug}", h.ProjectDetail)
	mux.HandleFunc("GET /lab", h.Lab)
	mux.HandleFunc("GET /journal", h.Journal)

	// API
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /api/status", h.Status)

	// Static assets
	fs := http.FileServer(http.Dir(staticDir))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	return mux
}

