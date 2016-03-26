package garson

import "net/http"

type Response struct {
    http.ResponseWriter
}

// NotFound returns a string decalring that the requested router
// is not found
func NotFound(res http.ResponseWriter) {
    res.WriteHeader(404)
    res.Write([]byte("Not Found"))
}

func (r *Response) Success(text string) {

}

func (r *Response) Error(err string) {

}


func (r *Response) JSON(status, obj map[string]interface{}) {

}

func (r *Response) Redirect(path string) {

}
