package garson

import (
	"net/http"
	"strings"
)

// Router struct
type Router struct {
	Trie             Trie
	beforeMiddleware []func(http.Handler) http.Handler
	afterMiddleware  []func(http.Handler) http.Handler
}

// add is a shortcut func to append new routes to the routes array
// and it extracts the params from the registered url
// used in router.Get(), router.Post(), router.Put(), router.Delete()
func (r *Router) add(method string, path string, handler http.HandlerFunc) {
	s := strings.Split(path, "/")
	r.Trie.Insert(method, s, handler)

}
