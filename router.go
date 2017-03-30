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

// Params is the type of paramters passed to a route
type Params map[string]string

// Route struct just contains the method, path and Handle func
type Route struct {
	Method           string
	Path             *regexp.Regexp
	RegisteredParams []string
	Handler          http.HandlerFunc
}

var paramsRegexp = regexp.MustCompile(`:(\w+)`)

// parseParams parses the request url against route.Path and returns
// a Params object
func parseParams(route *Route, path string) Params {
	matches := route.Path.FindAllStringSubmatch(path, -1)
	params := Params{}
	matchedParams := matches[0][1:]
	for k, v := range matchedParams {
		params[route.RegisteredParams[k]] = v
	}
	return params
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
// If the route is not found, it returns http.NotFound error
func (r *Router) Try(path string, method string) (http.HandlerFunc, Params, error) {
	for _, route := range r.Routes {
		if route.Method == method {
			match := route.Path.MatchString(path)
			if match == false {
				continue
			}
			params := Params{}
			// check if this route has registered params, and then parse them
			if len(route.RegisteredParams) > 0 {
				params = parseParams(route, path)
			}
			return route.Handler, params, nil
		}
	}
	return nil, Params{}, errors.New("Route not found")
}

// add is a shortcut func to append new routes to the routes array
// and it extracts the params from the registered url
// used in router.Get(), router.Post(), router.Put(), router.Delete()
func add(r *Router, method string, path string, handler http.HandlerFunc) {
	route := &Route{}
	route.Method = method
	path = "^" + path + "$"
	route.Handler = handler
	if strings.Contains(path, ":") {
		matches := paramsRegexp.FindAllStringSubmatch(path, -1)
		if matches != nil {
			for _, v := range matches {
				route.RegisteredParams = append(route.RegisteredParams, v[1])
				// remove the :params from the url path and replace them with regex
				path = strings.Replace(path, v[0], `(\w+)`, 1)
			}
		}
	}
	compiledPath, err := regexp.Compile(path)
	if err != nil {
		panic(err)
	}
	route.Path = compiledPath
	r.Routes = append(r.Routes, route)
}

// Get adds a GET method to routes
func (r *Router) Get(path string, handler http.HandlerFunc) {
	add(r, "GET", path, handler)
}

// Post adds a POST method to routes
func (r *Router) Post(path string, handler http.HandlerFunc) {
	add(r, "POST", path, handler)
}

// Put adds a PUT method to routes
func (r *Router) Put(path string, handler http.HandlerFunc) {
	add(r, "PUT", path, handler)
}

// Delete adds a DELETE method to routes
func (r *Router) Delete(path string, handler http.HandlerFunc) {
	add(r, "DELETE", path, handler)
}

// Head adds a HEAD method to routes
func (r *Router) Head(path string, handler http.HandlerFunc) {
	add(r, "HEAD", path, handler)
}

// Options adds an OPTIONS method to routes
func (r *Router) Options(path string, handler http.HandlerFunc) {
	add(r, "OPTIONS", path, handler)
}

// ServeHTTP implementats of the http.Handler interface
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, params, err := router.Try(r.URL.Path, r.Method)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	ctx := context.WithValue(r.Context(), "route_params", params)

	// execute the router handler
	handler(w, r.WithContext(ctx))
}
