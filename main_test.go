package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDefaultRouteCode(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()

	DefaultRoute(resp, req)

	if resp.Code != 200 {
		t.Errorf("GET %s: expected status code is %d; but got %d", req.URL, 200, resp.Code)
	}
}

func TestDefaultRouteBody(t *testing.T) {
	// If we have both positive/negative case for each API endpoint,
	// we can use SubTesting with t.Run()
	t.Run("Positive Case", func(t *testing.T) {
		expected := "{\"content\":\"hello, developers\"}\n"
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		resp := httptest.NewRecorder()

		// Call imported & targeted function
		DefaultRoute(resp, req)

		got := resp.Body.String()
		// fmt.Printf("type of got: %T \n", got)

		if got != expected {
			t.Errorf("got: [ %q ], expected: [ %q ] \n", got, expected)
			t.Errorf("Called endpoint: %s \n", req.URL)
			t.Errorf("Got Status code: %d \n", resp.Code)
		}
	})

}

func TestHealthCheckRouteCode(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	resp := httptest.NewRecorder()

	HealthCheckRoute(resp, req)

	if resp.Code != 200 {
		t.Errorf("GET %s: expected status code is %d; but got %d", req.URL, 200, resp.Code)
	}
}

func TestHealthCheckRouteBody(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	resp := httptest.NewRecorder()

	// Call imported & targeted function
	HealthCheckRoute(resp, req)

	got := resp.Body.String()
	expected := "{\"code\":200,\"status\":\"ok\"}\n"

	if got != expected {
		t.Errorf("got: [ %q ], expected: [ %q ] \n", got, expected)
		t.Errorf("Called endpoint: %s \n", req.URL)
		t.Errorf("Got Status code: %d \n", resp.Code)
	}
}
