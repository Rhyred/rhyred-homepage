package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Rhyred/rhyred-homepage/internal/services"
	"github.com/Rhyred/rhyred-homepage/web/templates/components"
	"github.com/Rhyred/rhyred-homepage/web/templates/pages"
)

// Handler holds dependencies for HTTP handlers.
type Handler struct {
	Data *services.DataService
}

// NewHandler creates a new Handler.
func NewHandler(data *services.DataService) *Handler {
	return &Handler{Data: data}
}

// Home renders the homepage.
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	profile := h.Data.GetProfile()
	projects := h.Data.GetFeaturedProjects()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := pages.Home(profile, projects).Render(r.Context(), w); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// About renders the about page.
func (h *Handler) About(w http.ResponseWriter, r *http.Request) {
	profile := h.Data.GetProfile()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := pages.About(profile).Render(r.Context(), w); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Projects renders the projects page or returns filtered project grid via HTMX.
func (h *Handler) Projects(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	isHTMX := r.Header.Get("HX-Request") == "true"

	var projectList = h.Data.GetProjectsByCategory(category)

	if isHTMX {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := components.ProjectGrid(projectList).Render(r.Context(), w); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	categories := h.Data.GetCategories()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := pages.Projects(projectList, categories, category).Render(r.Context(), w); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// ProjectDetail renders a single project page.
func (h *Handler) ProjectDetail(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	project, found := h.Data.GetProjectBySlug(slug)
	if !found {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := pages.ProjectDetail(project).Render(r.Context(), w); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Lab renders the digital laboratory page.
func (h *Handler) Lab(w http.ResponseWriter, r *http.Request) {
	profile := h.Data.GetProfile()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := pages.Lab(profile.LabAreas).Render(r.Context(), w); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Journal renders the journal page.
func (h *Handler) Journal(w http.ResponseWriter, r *http.Request) {
	profile := h.Data.GetProfile()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := pages.Journal(profile.Journal).Render(r.Context(), w); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Health returns a health check response.
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"time":   time.Now().UTC().Format(time.RFC3339),
	})
}

// Status returns the application status.
func (h *Handler) Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"app":     "RHYRED",
		"version": "0.1.0",
		"status":  "running",
		"time":    time.Now().UTC().Format(time.RFC3339),
	})
}
