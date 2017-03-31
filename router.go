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
func (r *Router) add(method string, path string, handler http.HandlerFunc) {
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

// Connect adds a CONNECT method to routes
func (r *Router) Connect(path string, handler http.HandlerFunc) {
	r.add("CONNECT", path, handler)
}

// Get adds a GET method to routes
func (r *Router) Get(path string, handler http.HandlerFunc) {
	r.add("GET", path, handler)
}

// Post adds a POST method to routes
func (r *Router) Post(path string, handler http.HandlerFunc) {
	r.add("POST", path, handler)
}

// Put adds a PUT method to routes
func (r *Router) Put(path string, handler http.HandlerFunc) {
	r.add("PUT", path, handler)
}

// Delete adds a DELETE method to routes
func (r *Router) Delete(path string, handler http.HandlerFunc) {
	r.add("DELETE", path, handler)
}

// Head adds a HEAD method to routes
func (r *Router) Head(path string, handler http.HandlerFunc) {
	r.add("HEAD", path, handler)
}

// Options adds an OPTIONS method to routes
func (r *Router) Options(path string, handler http.HandlerFunc) {
	r.add("OPTIONS", path, handler)
}

// ServeHTTP implementats of the http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, params, err := r.Try(req.URL.Path, req.Method)
	if err != nil {
		http.NotFound(w, req)
		return
	}

	ctx := context.WithValue(req.Context(), "route_params", params)

	// execute the router handler
	handler(w, req.WithContext(ctx))
}

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
