package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`abc`))
		return
	}))
	defer server.Close()
	length, err := pageLength(server.URL)
	if err != nil {
		t.Error("Expected error to equal nil")
	}
	if length != 3 {
		t.Errorf("Expected length to equal 3, got %d", length)
	}
}
