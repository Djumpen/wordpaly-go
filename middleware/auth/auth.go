package auth

import (
	"log"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("auth middleware", r.URL)
		next.ServeHTTP(w, r)
	})
}
