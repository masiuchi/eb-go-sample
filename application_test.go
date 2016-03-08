package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	ts := httptest.NewServer(router())
	defer ts.Close()

	r, err := http.Get(ts.URL + "/ping")
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}
	if r.StatusCode != 200 {
		t.Fatalf("Status code is not 200. %v", r.StatusCode)
	}
}
