package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		wantStatus     int
		wantBody       string
		wantContentType string
	}{
		{
			name:            "ルートパスでHello Worldを返す",
			path:            "/",
			wantStatus:      http.StatusOK,
			wantBody:        "Hello World",
			wantContentType: "text/html; charset=utf-8",
		},
		{
			name:            "任意のパスでHello Worldを返す",
			path:            "/foo/bar",
			wantStatus:      http.StatusOK,
			wantBody:        "Hello World",
			wantContentType: "text/html; charset=utf-8",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			w := httptest.NewRecorder()

			helloHandler(w, req)

			res := w.Result()
			if res.StatusCode != tt.wantStatus {
				t.Errorf("status = %d, want %d", res.StatusCode, tt.wantStatus)
			}
			if ct := res.Header.Get("Content-Type"); ct != tt.wantContentType {
				t.Errorf("Content-Type = %q, want %q", ct, tt.wantContentType)
			}
			body := w.Body.String()
			if !strings.Contains(body, tt.wantBody) {
				t.Errorf("body does not contain %q: %s", tt.wantBody, body)
			}
		})
	}
}
