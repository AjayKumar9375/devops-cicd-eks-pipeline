package server

import (
	"encoding/json"
	"net/http"
	"time"
)

type Config struct {
	AppName     string
	Environment string
	Version     string
	Port        string
	StartedAt   time.Time
}

type Server struct {
	config Config
}

type healthResponse struct {
	Status      string `json:"status"`
	AppName     string `json:"appName"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
	Uptime      string `json:"uptime"`
}

func New(cfg Config) Server {
	return Server{config: cfg}
}

func (s Server) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleIndex)
	mux.HandleFunc("/healthz", s.handleHealth)
	mux.HandleFunc("/readyz", s.handleReady)

	return mux
}

func (s Server) handleIndex(w http.ResponseWriter, _ *http.Request) {
	payload := map[string]string{
		"message":     "CI/CD pipeline demo service",
		"environment": s.config.Environment,
		"version":     s.config.Version,
	}

	writeJSON(w, http.StatusOK, payload)
}

func (s Server) handleHealth(w http.ResponseWriter, _ *http.Request) {
	payload := healthResponse{
		Status:      "ok",
		AppName:     s.config.AppName,
		Environment: s.config.Environment,
		Version:     s.config.Version,
		Uptime:      time.Since(s.config.StartedAt).Round(time.Second).String(),
	}

	writeJSON(w, http.StatusOK, payload)
}

func (s Server) handleReady(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ready"})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

