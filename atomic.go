package main

import (
	"crypto/tls"
	"net/http"
	"sync/atomic"
)

var client atomic.Value // stores *http.Client

// initializeClient sets the client if the certificate is available
func initializeClient(cert tls.Certificate) {
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	tr := &http.Transport{TLSClientConfig: tlsConfig}
	c := &http.Client{Transport: tr}
	client.Store(c)
}

// getClient safely retrieves the client if initialized
func getClient() *http.Client {
	v := client.Load()
	if v == nil {
		return nil
	}
	return v.(*http.Client)
}

// handler uses the client if available
func handler(w http.ResponseWriter, r *http.Request) {
	c := getClient()
	if c == nil {
		http.Error(w, "Client not initialized", http.StatusServiceUnavailable)
		return
	}

	// Use the client to make a request (example)
	resp, err := c.Get("https://example.com")
	if err != nil {
		http.Error(w, "Failed to make request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request successful"))
}
