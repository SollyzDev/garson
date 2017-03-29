package garson

import (
	"context"
	"errors"
	"net/http"
	"regexp"
	"strings"
)

// Router struct
type Router struct {
	Routes []*Route
}

// Route struct just contains the method, path and Handle func
type Route struct {
	Method           string
	Path             string
	RegisteredParams []string
	Params           map[string]string
	Handler          http.HandlerFunc
}

func (r *Route) parseParams(re *regexp.Regexp, path string) {
	matches := re.FindAllStringSubmatch(path, -1)
	r.Params = make(map[string]string)
	params := matches[0][1:len(matches[0])]
	for k, v := range params {
		r.Params[r.RegisteredParams[k]] = v
	}
}

// New creates and return a new router object
// it should be passed to http.ListenAndServe
// example:
// router := garson.New()
// router.Get("/hello", func(w http.ResponseWriter, r *http.Request){})
// http.ListenAndServe(":8080", router)
func New() *Router {
	return &Router{}
}

// Try loops through the routes array to find the requested route
// If the route is not found, it returns NotFound error
// FIX: should return (*Router,error)
func (r *Router) Try(path string, method string) (*Route, error) {
	for _, route := range r.Routes {
		if route.Method == method {
			re := regexp.MustCompile(route.Path)
			match := re.MatchString(path)
			// check if the registered route has params
			if match {
				if len(route.RegisteredParams) > 0 {
					route.parseParams(re, path)
				}
				return route, nil
			}
		}
	}
	return &Route{}, errors.New("Route not found")
}

// NotFound returns a string decalring that the requested route
// was not found
func NotFound(w http.ResponseWriter) {
	w.WriteHeader(404)
	w.Write([]byte("Not Found"))
}

// add is a shortcut func to append new routes to the routes array
// and it extracts the params from the registered url
// used in router.Get(), router.Post(), router.Put(), router.Delete()
func add(r *Router, method string, path string, handler http.HandlerFunc) {
	route := &Route{}
	route.Method = method
	route.Path = "^" + path + "$"
	route.Handler = handler
	if strings.Contains(route.Path, ":") {
		re := regexp.MustCompile(`:(\w+)`)
		matches := re.FindAllStringSubmatch(route.Path, -1)
		if matches != nil {
			for _, v := range matches {
				route.RegisteredParams = append(route.RegisteredParams, v[1])
				// remove the :params from the url path and replace them with regex
				route.Path = strings.Replace(route.Path, v[0], `(\w+)`, 1)
			}
		}
	}
	r.Routes = append(r.Routes, route)
}

// Get the adds a GET method to routes
func (r *Router) Get(path string, handler http.HandlerFunc) {
	add(r, "GET", path, handler)
}

// Post the adds a POST method to routes
func (r *Router) Post(path string, handler http.HandlerFunc) {
	add(r, "POST", path, handler)
}

// Put the adds a PUT method to routes
func (r *Router) Put(path string, handler http.HandlerFunc) {
	add(r, "PUT", path, handler)
}

// Delete the adds a DELETE method to routes
func (r *Router) Delete(path string, handler http.HandlerFunc) {
	add(r, "DELETE", path, handler)
}

// ServeHTTP implementats of the http.Handler interface
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	route, err := router.Try(r.URL.Path, r.Method)
	if err != nil {
		NotFound(w)
		return
	}
	ctx := context.WithValue(r.Context(), "route_params", route.Params)

	route.Handler(w, r.WithContext(ctx))
}
