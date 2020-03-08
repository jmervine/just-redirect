package main

import (
    "os"
    "testing"
    "net/http"
    "net/http/httptest"
)

func TestRedirect(t *testing.T) {
    targetServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
    defer targetServer.Close()

    target = targetServer.URL
    server := httptest.NewServer(http.HandlerFunc(redirect))
    defer server.Close()

    resp, err := http.Get(server.URL)
    if err != nil {
        t.Errorf("Expected nil, got: %v", err)
    }

    if resp.StatusCode != 200 {
        t.Errorf("Expected 200, got: %d", resp.StatusCode)
    }

    req := resp.Request
    if req.Method != "GET" {
        t.Errorf("Expected GET, got: %s", resp.Request.Method)
    }

    if req.URL.String() != targetServer.URL {
        t.Errorf("Expected %s, got: %s", targetServer.URL, resp.Request.URL.RequestURI())
    }
}

func TestConfig(t *testing.T) {
    tgt := "https://www.example.com"
    os.Setenv("REDIRECT_TARGET", tgt)
    os.Setenv("PORT", "5432")
    os.Setenv("BIND", "localhost")

    defer func() {
        _ = os.Unsetenv("REDIRECT_TARGET")
        _ = os.Unsetenv("PORT")
        _ = os.Unsetenv("BIND")
    }()

    config()

    if target != tgt {
        t.Errorf("Expected %s, got: %s", tgt, target)
    }

    if port != "5432" {
        t.Errorf("Expected 5432, got: %s", port)
    }

    if bind != "localhost" {
        t.Errorf("Expected localhost, got: %s", bind)
    }
}

func TestConfig_defaults(t *testing.T) {
    tgt := "https://www.example.com"
    os.Setenv("REDIRECT_TARGET", tgt)

    defer func() {
        _ = os.Unsetenv("REDIRECT_TARGET")
    }()

    config()

    if target != tgt {
        t.Errorf("Expected %s, got: %s", tgt, target)
    }

    if port != "3000" {
        t.Errorf("Expected 3000, got: %s", port)
    }

    if bind != "0.0.0.0" {
        t.Errorf("Expected 0.0.0.0, got: %s", bind)
    }
}
