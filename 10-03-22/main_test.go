package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io"
)

func TestServeHTPP(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/yey", nil)
	w := httptest.NewRecorder()

	ServeHTPP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code: %#v, got: %#v.", http.StatusOK, w.Code)
	}

	result := w.Result()
	body, err := io.ReadAll(result.Body)

	if err != nil {
		t.Fatal(err)
	}

	if result.Header.Get("YOY") != "Cool" {
		t.Fatalf("Expected a header of: %#v, got: %#v", "Cool", result.Header.Get("YOY"))
	}

	if string(body) != "Server is up." {
		t.Fatalf("Expected body of: %#v. got: %#v.", "Server is up", string(body))
	}
	
}