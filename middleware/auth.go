package middleware

import (
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Authentication testing")

		next.ServeHTTP(rw, r)
	})
}
