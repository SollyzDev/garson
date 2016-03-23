package main

import "net/http"

type Router struct {
    Routes []Route
}

type Route struct {
    Method string
    Path string
    Handler HandlerFunc
}

type HandlerFunc func(res http.ResponseWriter, req *http.Request)

func (r Router) Get(path string, handler HandlerFunc) {
    route := Route{}
    route.Method = "GET"
    route.Path = path
    route.Handler = handler
    r.Routes = append(r.Routes, route)
}

func (r Router) Post(route string) {

}

func (r Router) Put(route string) {

}

func (r Router) Delete(route string) {

}
