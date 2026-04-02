package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHealthEndpointIncludesAppMetadata(t *testing.T) {
	srv := New(Config{
		AppName:     "demo-service",
		Environment: "dev",
		Version:     "abc1234",
		StartedAt:   time.Now().Add(-2 * time.Minute),
	})

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	srv.Routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()
	for _, token := range []string{"demo-service", "dev", "abc1234", "ok"} {
		if !strings.Contains(body, token) {
			t.Fatalf("expected response body to contain %q, got %s", token, body)
		}
	}
}

func TestReadinessEndpoint(t *testing.T) {
	srv := New(Config{StartedAt: time.Now()})

	req := httptest.NewRequest(http.MethodGet, "/readyz", nil)
	rec := httptest.NewRecorder()

	srv.Routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "ready") {
		t.Fatalf("expected readiness response, got %s", rec.Body.String())
	}
}

