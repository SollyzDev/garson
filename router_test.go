package garson

import (
	"net/http"
	"testing"
)

func TestBasicRouter(t *testing.T) {
	router := New()
	router.Get("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
}
