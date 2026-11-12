package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Rhyred/rhyred-homepage/internal/handlers"
	"github.com/Rhyred/rhyred-homepage/internal/services"
)

func setupTestRouter(t *testing.T) http.Handler {
	t.Helper()
	dataSvc := services.NewDataService("../../web/data")
	if err := dataSvc.Load(); err != nil {
		t.Fatalf("failed to load data: %v", err)
	}
	h := handlers.NewHandler(dataSvc)
	return New(h, "../../web/static")
}

func TestSecurityHeaders(t *testing.T) {
	r := setupTestRouter(t)
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	headers := map[string]string{
		"X-Content-Type-Options": "nosniff",
		"X-Frame-Options":        "DENY",
		"Referrer-Policy":        "strict-origin-when-cross-origin",
	}

	for k, expected := range headers {
		if val := w.Header().Get(k); val != expected {
			t.Errorf("header %s: expected %q, got %q", k, expected, val)
		}
	}

	if csp := w.Header().Get("Content-Security-Policy"); csp == "" {
		t.Error("expected Content-Security-Policy header to be present")
	}
}

func TestFaviconRoute(t *testing.T) {
	r := setupTestRouter(t)
	req := httptest.NewRequest("GET", "/favicon.ico", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200 for /favicon.ico, got %d", w.Code)
	}
	if ct := w.Header().Get("Content-Type"); ct != "image/svg+xml" {
		t.Errorf("expected image/svg+xml, got %s", ct)
	}
	if cc := w.Header().Get("Cache-Control"); cc != "public, max-age=86400" {
		t.Errorf("expected public, max-age=86400, got %s", cc)
	}
}

func TestStaticAssetsCaching(t *testing.T) {
	r := setupTestRouter(t)
	req := httptest.NewRequest("GET", "/static/css/variables.css", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200 for static asset, got %d", w.Code)
	}
	if cc := w.Header().Get("Cache-Control"); cc != "public, max-age=86400" {
		t.Errorf("expected Cache-Control header, got %s", cc)
	}
}

func TestStaticDirectoryListingDisabled(t *testing.T) {
	r := setupTestRouter(t)

	testCases := []string{
		"/static/",
		"/static/css/",
		"/static/js/",
	}

	for _, tc := range testCases {
		req := httptest.NewRequest("GET", tc, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected 404 for directory %s, got %d", tc, w.Code)
		}
	}
}

