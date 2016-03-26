package garson

import "net/http"

type Router struct {
    Routes []Route
}

type Route struct {
    Method string
    Path string
    Handler HandlerFunc
}

type HandlerFunc func(req *http.Request, res http.ResponseWriter)

func New() Router {
    return Router{}
}

func (r *Router) Try(path string, method string) (*Route,string) { // FIX: should return (*Router,error)
    for _, route := range r.Routes {
        if route.Method == method && route.Path == path {
            return &route, ""
        }
    }
    return &Route{}, "route not found"
}

func (r *Router) Get(path string, handler HandlerFunc) {
    route := Route{}
    route.Method = "GET"
    route.Path = path
    route.Handler = handler
    r.Routes = append(r.Routes, route)
}

func (r *Router) Post(route string) {

}

func (r *Router) Put(route string) {

}

func (r *Router) Delete(route string) {

}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    route, err :=  r.Try(req.URL.Path, req.Method)
    if err != "" {
        panic(err)
    }
    route.Handler(req, w)
}