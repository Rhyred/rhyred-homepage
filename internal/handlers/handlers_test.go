package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Rhyred/rhyred-homepage/internal/services"
)

func setupTestHandler(t *testing.T) *Handler {
	t.Helper()
	dataSvc := services.NewDataService("../../web/data")
	if err := dataSvc.Load(); err != nil {
		t.Fatalf("failed to load test data: %v", err)
	}
	return NewHandler(dataSvc)
}

func TestHealth(t *testing.T) {
	h := setupTestHandler(t)
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	h.Health(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var result map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}
	if result["status"] != "ok" {
		t.Errorf("expected status ok, got %s", result["status"])
	}
}

func TestStatus(t *testing.T) {
	h := setupTestHandler(t)
	req := httptest.NewRequest("GET", "/api/status", nil)
	w := httptest.NewRecorder()

	h.Status(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var result map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}
	if result["app"] != "RHYRED" {
		t.Errorf("expected app RHYRED, got %v", result["app"])
	}
	if result["status"] != "running" {
		t.Errorf("expected status running, got %v", result["status"])
	}
}

func TestHome(t *testing.T) {
	h := setupTestHandler(t)
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	h.Home(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "text/html; charset=utf-8" {
		t.Errorf("expected content type text/html, got %s", w.Header().Get("Content-Type"))
	}
	body := w.Body.String()
	if len(body) < 100 {
		t.Error("expected substantial HTML response")
	}
}

func TestHome404(t *testing.T) {
	h := setupTestHandler(t)
	req := httptest.NewRequest("GET", "/nonexistent", nil)
	w := httptest.NewRecorder()

	h.Home(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", w.Code)
	}
}

func TestProjects(t *testing.T) {
	h := setupTestHandler(t)
	req := httptest.NewRequest("GET", "/projects", nil)
	w := httptest.NewRecorder()

	h.Projects(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestProjectsHTMXFilter(t *testing.T) {
	h := setupTestHandler(t)
	req := httptest.NewRequest("GET", "/projects?category=Web+Application", nil)
	req.Header.Set("HX-Request", "true")
	w := httptest.NewRecorder()

	h.Projects(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	// HTMX responses should be partial HTML (no full page)
	body := w.Body.String()
	if len(body) == 0 {
		t.Error("expected non-empty response for HTMX request")
	}
}

func TestProjectDetailNotFound(t *testing.T) {
	h := setupTestHandler(t)
	req := httptest.NewRequest("GET", "/projects/nonexistent", nil)
	req.SetPathValue("slug", "nonexistent")
	w := httptest.NewRecorder()

	h.ProjectDetail(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", w.Code)
	}
}

func TestProjectDetailFound(t *testing.T) {
	h := setupTestHandler(t)
	req := httptest.NewRequest("GET", "/projects/satset", nil)
	req.SetPathValue("slug", "satset")
	w := httptest.NewRecorder()

	h.ProjectDetail(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestAbout(t *testing.T) {
	h := setupTestHandler(t)
	req := httptest.NewRequest("GET", "/about", nil)
	w := httptest.NewRecorder()

	h.About(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestLab(t *testing.T) {
	h := setupTestHandler(t)
	req := httptest.NewRequest("GET", "/lab", nil)
	w := httptest.NewRecorder()

	h.Lab(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestJournal(t *testing.T) {
	h := setupTestHandler(t)
	req := httptest.NewRequest("GET", "/journal", nil)
	w := httptest.NewRecorder()

	h.Journal(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}
