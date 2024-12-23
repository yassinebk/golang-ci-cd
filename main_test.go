package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "valid path",
			path:           "/",
			expectedStatus: http.StatusOK,
			expectedBody:   "This request was processed by host",
		},
		{
			name:           "invalid path",
			path:           "/invalid",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "404 page not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			w := httptest.NewRecorder()

			homeHandler(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, w.Code)
			}

			if !strings.Contains(w.Body.String(), tt.expectedBody) {
				t.Errorf("Expected body to contain %q, got %q", tt.expectedBody, w.Body.String())
			}
		})
	}
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	healthHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response struct {
		Status string `json:"status"`
		Uptime string `json:"uptime"`
	}

	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Status != "healthy" {
		t.Errorf("Expected status 'healthy', got %q", response.Status)
	}

	if response.Uptime == "" {
		t.Error("Expected non-empty uptime")
	}
}

func TestInfoHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/info", nil)
	w := httptest.NewRecorder()

	infoHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response ServerInfo
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Hostname == "" {
		t.Error("Expected non-empty hostname")
	}

	if response.TimeStarted.IsZero() {
		t.Error("Expected non-zero start time")
	}
}

func TestMetricsHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/metrics", nil)
	w := httptest.NewRecorder()

	metricsHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response struct {
		MemoryAlloc uint64 `json:"memoryAlloc"`
		Goroutines  int    `json:"goroutines"`
	}

	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.MemoryAlloc != 0 {
		t.Errorf("Expected MemoryAlloc 0, got %d", response.MemoryAlloc)
	}

	if response.Goroutines != 0 {
		t.Errorf("Expected Goroutines 0, got %d", response.Goroutines)
	}
}

func TestMainFunction(t *testing.T) {
	go func() {
		main()
	}()

	resp, err := http.Get("http://localhost:80/health")
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
