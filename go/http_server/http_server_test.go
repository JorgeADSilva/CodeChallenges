package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Home)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Home handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "text/html; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Home handler returned wrong content type: got %v want %v", contentType, expectedContentType)
	}
}

func TestInfoHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/info", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Info)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Info handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedBody := "Info Page"
	if body := rr.Body.String(); body != expectedBody {
		t.Errorf("Info handler returned unexpected body: got %v want %v", body, expectedBody)
	}
}
