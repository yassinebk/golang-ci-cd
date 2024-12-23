package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// ServerInfo contains basic information about the server
type ServerInfo struct {
	Hostname    string    `json:"hostname"`
	TimeStarted time.Time `json:"timeStarted"`
}

var (
	startTime  = time.Now()
	serverInfo ServerInfo
)

func init() {
	hostname, _ := os.Hostname()
	serverInfo = ServerInfo{
		Hostname:    hostname,
		TimeStarted: startTime,
	}
}

// homeHandler displays the hostname
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "<h1>This request was processed by host: %s</h1>\n", serverInfo.Hostname)
}

// healthHandler returns server health status
func healthHandler(w http.ResponseWriter, r *http.Request) {
	health := struct {
		Status string `json:"status"`
		Uptime string `json:"uptime"`
	}{
		Status: "healthy",
		Uptime: time.Since(startTime).String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}

// infoHandler returns detailed server information
func infoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(serverInfo)
	if err != nil {
		w.Write([]byte("error"))
	}
}

// metricsHandler returns basic metrics
func metricsHandler(w http.ResponseWriter, r *http.Request) {
	metrics := struct {
		MemoryAlloc uint64 `json:"memoryAlloc"`
		Goroutines  int    `json:"goroutines"`
	}{
		MemoryAlloc: 0, // You would typically get this from runtime.MemStats
		Goroutines:  0, // You would typically get this from runtime.NumGoroutine
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/info", infoHandler)
	mux.HandleFunc("/metrics", metricsHandler)

	// Start the server
	fmt.Fprintf(os.Stdout, "Web Server started. Listening on 0.0.0.0:80\n")
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting server: %v\n", err)
		os.Exit(1)
	}
}
