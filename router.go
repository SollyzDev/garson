package garson

import (
	"errors"
	"net/http"
	"regexp"
	"strings"
)

// Router struct
type Router struct {
	Routes []Route
}

// Route struct just contains the method, path and Handle func
type Route struct {
	Method           string
	Path             string
	RegisteredParams []string
	Params           map[string]string
	Handler          Handle
}

// Handle serves as a type of the func that is qoing to be fired
// when the routers finds the requested route
type Handle func(ctx *Context)

// New creates and return a new router object
// it should be passed to http.ListenAndServe
// example:
// router := garson.New()
// router.Get("/hello", func(ctx *garson.Context){})
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

				matches := re.FindAllStringSubmatch(path, -1)
				route.Params = make(map[string]string)
				params := matches[0][1:len(matches[0])]
				for k, v := range params {
					route.Params[route.RegisteredParams[k]] = v
				}
				return &route, nil
			}
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
// and it extracts the params from the registered url
// used in router.Get(), router.Post(), router.Put(), router.Delete()
func add(r *Router, method string, path string, handler Handle) {
	route := Route{}
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
	ctx := NewContext(route, res, req)
	route.Handler(ctx)
}
