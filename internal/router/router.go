package router

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

	// Favicon shortcut
	mux.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		faviconPath := filepath.Join(staticDir, "images", "favicon.svg")
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Header().Set("Cache-Control", "public, max-age=86400")
		http.ServeFile(w, r, faviconPath)
	})

	// Static assets with cache-control & directory listing disabled
	mux.HandleFunc("GET /static/", func(w http.ResponseWriter, r *http.Request) {
		relPath := strings.TrimPrefix(r.URL.Path, "/static/")
		if relPath == "" || strings.HasSuffix(relPath, "/") {
			http.NotFound(w, r)
			return
		}

		fullPath := filepath.Join(staticDir, filepath.Clean("/"+relPath))
		info, err := os.Stat(fullPath)
		if err != nil || info.IsDir() {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Cache-Control", "public, max-age=86400")
		http.ServeFile(w, r, fullPath)
	})

	return securityHeadersMiddleware(mux)
}

func securityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		w.Header().Set("Permissions-Policy", "geolocation=(), camera=(), microphone=()")
		w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline' https://unpkg.com; style-src 'self' 'unsafe-inline' https://fonts.googleapis.com; font-src 'self' https://fonts.gstatic.com; img-src 'self' data:; connect-src 'self';")

		next.ServeHTTP(w, r)
	})
}
