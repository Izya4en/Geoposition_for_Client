package test

import (
	"net/http"
	"net/http/httptest"
	"tedx-maps/internal/server"
	"testing"
)

func TestAPIHealthCheck(t *testing.T) {
	s := server.New()
	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	w := httptest.NewRecorder()

	s.Router().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}
