package garson

import (
	"encoding/json"
	// "fmt"
	"net/http"
)

type Context struct {
	Request     *http.Request
	Response    http.ResponseWriter
	QueryParams map[string]string
	RouteParams map[string]string
	BodyParams  map[string]string
}

func NewContext(route *Route, res http.ResponseWriter, req *http.Request) *Context {
	ctx := Context{}
	ctx.Request = req
	ctx.Response = res
	ctx.RouteParams = route.Params
	return &ctx
}

// GetParam gets the value of a param, if the param
// is not found, the default_val will bre returned
func (ctx *Context) GetParam(key string, default_val interface{}) string {
	val, exists := ctx.RouteParams[key]
	if exists {
		return val
	}
	val, exists = ctx.QueryParams[key]
	if exists {
		return val
	}
	val, exists = ctx.BodyParams[key]
	if exists {
		return val
	}
	return default_val.(string)
}

// NotFound returns a 404 response
func (ctx *Context) NotFound() {
	ctx.Response.WriteHeader(404)
	ctx.Response.Write([]byte("Not Found"))
}

func (ctx *Context) Write(text string) {
	header := ctx.Response.Header()
	header.Set("Content-Type", "text/plain")
	ctx.Response.Write([]byte(text))
}

func (ctx *Context) Success(text string) {

}

func (ctx *Context) Error(err string) {

}

func (ctx *Context) Json(obj map[string]interface{}) {
	header := ctx.Response.Header()
	header.Set("Content-Type", "application/json")
	enc, err := json.Marshal(obj)
	if err != nil {
		panic("couldnt encode response body to json")
	}
	ctx.Response.Write(enc)
}

func (ctx *Context) Redirect(path string) {

}
