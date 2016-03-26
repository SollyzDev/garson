package garson

import "net/http"

// Router struct
type Router struct {
    Routes []Route
}

// Route struct just contains the method, path and Handle func
type Route struct {
    Method string
    Path string
    Handler Handle
}

// Handle serves as a type of the func that is qoing to be fired 
// when the routers finds the requested route
type Handle func(req *http.Request, res http.ResponseWriter)

// New creates and return a new router object
// it should be passed to http.ListenAndServe
// example:
// router := garson.New()
// router.Get("/hello", func(req *http.Request, res http.ResponseWriter){})
// http.ListenAndServe(":8080", router)
func New() *Router {
    return &Router{}
}

// Try loops through the routes array to find the requested route
// If the route is not found, it returns NotFound error
// FIX: should return (*Router,error)
func (r *Router) Try(path string, method string) (*Route,string) { 
    for _, route := range r.Routes {
        if route.Method == method && route.Path == path {
            return &route, ""
        }
    }
    return &Route{}, "route not found"
}

// a shortcut func to append new routes to the routes array
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
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    route, err :=  r.Try(req.URL.Path, req.Method)
    if err != "" {
        //panic(err)
        return
    }
    route.Handler(req, w)
}