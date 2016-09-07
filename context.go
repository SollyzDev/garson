package garson

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Request     *http.Request
	Response    http.ResponseWriter
	QueryParams map[string]interface{}
	RouteParams map[string]interface{}
	BodyParams  map[string]interface{}
}

func NewContext(res http.ResponseWriter, req *http.Request) *Context {
	ctx := Context{}
	ctx.Request = req
	ctx.Response = res
	return &ctx
}

// NotFound returns a 404 response
func (ctx *Context) NotFound() {
	ctx.Response.WriteHeader(404)
	ctx.Response.Write([]byte("Not Found"))
}

func (ctx *Context) Write(text string) {
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
