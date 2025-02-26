package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRestRoute(t *testing.T) {
	t.Run("Check status code", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/", nil)
		resp := httptest.NewRecorder()

		RestRoute(resp, req)

		if resp.Code != 200 {
			t.Errorf("GET %s: expected status code is %d; but got %d", req.URL, 200, resp.Code)
		}
	})

	t.Run("Check response body", func(t *testing.T) {
		expected := "{\"content\":\"API root!\"}\n"

		req, _ := http.NewRequest(http.MethodGet, "/api/v1/", nil)
		resp := httptest.NewRecorder()

		// Call imported & targeted function
		RestRoute(resp, req)

		got := resp.Body.String()

		if got != expected {
			t.Errorf("got: [ %q ], expected: [ %q ] \n", got, expected)
			t.Errorf("Called endpoint: %s \n", req.URL)
			t.Errorf("Got Status code: %d \n", resp.Code)
		}
	})
}

func TestGetRandomUuid(t *testing.T) {
	t.Run("Check status code", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/uuid", nil)
		resp := httptest.NewRecorder()

		GetRandomUuid(resp, req)

		if resp.Code != 200 {
			t.Errorf("GET %s: expected status code is %d; but got %d", req.URL, 200, resp.Code)
		}
	})
}
