package garson

import (
	"errors"
	"net/http"
)

// Router struct
type Router struct {
	Routes []Route
}

// Route struct just contains the method, path and Handle func
type Route struct {
	Method  string
	Path    string
	Handler Handle
}

// Handle serves as a type of the func that is qoing to be fired
// when the routers finds the requested route
type Handle func(ctx *Context)

// New creates and return a new router object
// it should be passed to http.ListenAndServe
// example:
// router := garson.New()
// router.Get("/hello", func(res http.ResponseWriter, req *http.Request){})
// http.ListenAndServe(":8080", router)
func New() *Router {
	return &Router{}
}

// Try loops through the routes array to find the requested route
// If the route is not found, it returns NotFound error
// FIX: should return (*Router,error)
func (r *Router) Try(path string, method string) (*Route, error) {
	for _, route := range r.Routes {
		if route.Method == method && route.Path == path {
			return &route, nil
		}
	}
	return &Route{}, errors.New("Route not found")
}

// NotFound returns a string decalring that the requested router
// is not found
func NotFound(res http.ResponseWriter) {
	res.WriteHeader(404)
	res.Write([]byte("Not Found"))
}

// add is a shortcut func to append new routes to the routes array
// used in router.Get(), router.Post(), router.Put(), router.Delete()
func add(r *Router, method string, path string, handler Handle) {
	route := Route{}
	route.Method = method
	route.Path = path
	route.Handler = handler
	r.Routes = append(r.Routes, route)
}

// Get the adds a GET method to routes
func (r *Router) Get(path string, handler Handle) {
	add(r, "GET", path, handler)
}

// Post the adds a POST method to routes
func (r *Router) Post(path string, handler Handle) {
	add(r, "POST", path, handler)
}

// Put the adds a PUT method to routes
func (r *Router) Put(path string, handler Handle) {
	add(r, "PUT", path, handler)
}

// Delete the adds a DELETE method to routes
func (r *Router) Delete(path string, handler Handle) {
	add(r, "DELETE", path, handler)
}

// ServeHTTP implementats of the http.Handler interface
func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	route, err := r.Try(req.URL.Path, req.Method)
	if err != nil {
		NotFound(res)
		return
	}
	ctx := NewContext(res, req)
	route.Handler(ctx)
}
