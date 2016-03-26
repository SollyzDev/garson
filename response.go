package garson

import "net/http"

type Response struct {
    http.ResponseWriter
}

func (r Response) Success(text string) {

}

func (r Response) Error(err string) {

}


func (r Response) JSON(status, obj map[string]interface{}) {

}

func (r Response) Redirect(path string) {

}
