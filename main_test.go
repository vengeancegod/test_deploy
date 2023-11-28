package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDockerHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(dockerHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code")
	}

	expected := "hello from dockerasdasdsadsa"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body")
	}
}
