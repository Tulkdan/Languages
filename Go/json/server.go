package main

import (
	"crypto/subtle"
	"net/http"
)

type BasicAuthCredentials struct {
	User     string
	Password string
}

func DecorateWithBasicAuth(next http.HandlerFunc, credentials *BasicAuthCredentials) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, password, ok := r.BasicAuth()

		const noMatch = 0
		if !ok || user != credentials.User || subtle.ConstantTimeCompare([]byte(credentials.Password), []byte(password)) == noMatch {
			w.Header().Set("WWW-Authenticate", `Basic realm="Resctricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid credentials"))
			return
		}

		next.ServeHTTP(w, r)
	}
}
