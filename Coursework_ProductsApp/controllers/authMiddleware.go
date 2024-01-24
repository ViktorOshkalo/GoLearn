package controllers

import (
	"main/configuration"
	"net/http"
)

func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, success := r.BasicAuth()
		if !success {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if username != configuration.User {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if password != configuration.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
