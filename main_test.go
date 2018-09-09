package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNormalRequest(t *testing.T) {
	request, _ := http.NewRequest("GET", "/view/hello", nil)
	response := httptest.NewRecorder()

	viewHandler(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Non-expected status code%v:\n\tbody: %v", "200", response.Code)
	}
}
