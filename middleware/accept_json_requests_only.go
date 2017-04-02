package middleware

import (
	"net/http"
	"strings"
)

func AcceptJSONRequestsOnly(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctHeader := r.Header.Get("Content-Type")
		ct := strings.Split(ctHeader, ";")[0]
		if ct != "application/json" {
			http.Error(w, "Only JSON requests are allowed", http.StatusBadRequest)
		} else {
			// don't execute the next handler unless the request is JSON
			next.ServeHTTP(w, r)
		}

	}
	return http.HandlerFunc(fn)
}
