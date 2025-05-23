package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_AuthWithValidPassword_Gives200(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello world!</body></html>")
	}
	w := httptest.NewRecorder()

	wantUser := "admin"
	wantPassword := "password"
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	r.SetBasicAuth(wantUser, wantPassword)

	wantCredentials := &BasicAuthCredentials{
		User:     wantUser,
		Password: wantPassword,
	}

	decorated := DecorateWithBasicAuth(handler, wantCredentials)
	decorated.ServeHTTP(w, r)

	wantCode := http.StatusOK

	if w.Code != wantCode {
		t.Fatalf("status code, want %d, got %d", wantCode, w.Code)
	}

	gotAuth := w.Header().Get("WWW-Authenticate")
	wantAuth := ``
	if gotAuth != wantAuth {
		t.Fatalf("WWW-Authenticate, want: %s, got: %s", wantAuth, gotAuth)
	}
}

func Test_AUthWithInvalidPassword_Gives403(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<tml><body>Hello world!</body></html>")
	}
	w := httptest.NewRecorder()

	wantUser := "admin"
	wantPassword := "password"
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	r.SetBasicAuth(wantUser, wantPassword)

	wantCredentials := &BasicAuthCredentials{
		User:     wantUser,
		Password: "",
	}

	decorated := DecorateWithBasicAuth(handler, wantCredentials)
	decorated.ServeHTTP(w, r)

	wantCode := http.StatusUnauthorized

	if w.Code != wantCode {
		t.Fatalf("status code, want %d, got %d", wantCode, w.Code)
	}

	gotAuth := w.Header().Get("WWW-Authenticate")
	wantAuth := `Basic realm="Resctricted"`
	if gotAuth != wantAuth {
		t.Fatalf("WWW-Authenticate, want: %s, got: %s", wantAuth, gotAuth)
	}
}
