package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func TestGetClient_NotInitialized(t *testing.T) {
	// Reset the client to nil for this test
	client = atomic.Value{}

	got := getClient()
	if got != nil {
		t.Errorf("Expected nil, got %v", got)
	}
}

func TestGetClient_Initialized(t *testing.T) {
	// Reset the client to nil for this test
	client = atomic.Value{}

	expected := &http.Client{}
	client.Store(expected)

	got := getClient()
	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

// mockRoundTripper implements http.RoundTripper for testing
type mockRoundTripper struct {
	resp *http.Response
	err  error
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.resp, m.err
}

func TestHandler_ClientNotInitialized(t *testing.T) {
	// Reset the client to nil
	client = atomic.Value{}

	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	handler(rr, req)

	if rr.Code != http.StatusServiceUnavailable {
		t.Errorf("Expected status %d, got %d", http.StatusServiceUnavailable, rr.Code)
	}
	if !bytes.Contains(rr.Body.Bytes(), []byte("Client not initialized")) {
		t.Errorf("Expected error message in response body")
	}
}

func TestHandler_ClientInitialized_RequestSuccess(t *testing.T) {
	// Prepare a mock client that returns a successful response
	client = atomic.Value{}
	mockResp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBufferString("OK")),
	}
	mockClient := &http.Client{
		Transport: &mockRoundTripper{resp: mockResp, err: nil},
	}
	client.Store(mockClient)

	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	handler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, rr.Code)
	}
	if !bytes.Contains(rr.Body.Bytes(), []byte("Request successful")) {
		t.Errorf("Expected success message in response body")
	}
}

func TestHandler_ClientInitialized_RequestFails(t *testing.T) {
	// Prepare a mock client that returns an error
	client = atomic.Value{}
	mockClient := &http.Client{
		Transport: &mockRoundTripper{resp: nil, err: http.ErrHandlerTimeout},
	}
	client.Store(mockClient)

	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	handler(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status %d, got %d", http.StatusInternalServerError, rr.Code)
	}
	if !bytes.Contains(rr.Body.Bytes(), []byte("Failed to make request")) {
		t.Errorf("Expected failure message in response body")
	}
}
