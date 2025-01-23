package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDefaultRoute(t *testing.T) {
	// expected := "{\"content\":\"hello, developers\"}"
	mock := httptest.NewServer(http.HandlerFunc(DefaultRoute))
	defer mock.Close()

	resp, err := http.Get(mock.URL + "/")
	if err != nil {
		t.Error(err)
	}

	txt, err := io.ReadAll(resp.Body)
	fmt.Printf("Got: %s", string(txt))

	// txt, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	t.Errorf("Failed to read response body")
	// 	t.Error(err)
	// }
	resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("GET %s: expected status code is %d; got %d", mock.URL, 200, resp.StatusCode)
	}
	// if string(txt) != expected {
	// 	t.Errorf("expected body [ %v ]; got [ %v ]", expected, string(txt))
	// }
}

func TestHealthCheckRoute(t *testing.T) {
	// // expected := "{\"code\":200,\"status\":\"ok\"}"
	// expected := []byte(`{"code":200,"status":"ok"}`)
	mock := httptest.NewServer(http.HandlerFunc(HealthCheckRoute))
	defer mock.Close()

	resp, err := http.Get(mock.URL + "/health")
	if err != nil {
		t.Error(err)
	}

	txt, err := io.ReadAll(resp.Body)
	fmt.Printf("Got: %s", string(txt))

	// txt, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	t.Errorf("Failed to read response body")
	// 	t.Error(err)
	// }
	resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("GET %s: expected status code is %d; got %d", mock.URL, 200, resp.StatusCode)
	}
	// if string(txt) != string(expected) {
	// 	fmt.Printf("%v\n", expected)
	// 	fmt.Printf("%v\n", txt)
	// 	t.Errorf("expected body [ %v ]; got [ %v ]", expected, txt)
	// }
}
