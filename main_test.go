package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNormalRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(viewHandler))
	defer ts.Close()
}
