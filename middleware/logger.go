package middleware

import (
	"log"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
