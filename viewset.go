package garson

import "net/http"

type ViewSet interface {
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	PutSingle(w http.ResponseWriter, r *http.Request)
	DeleteSingle(w http.ResponseWriter, r *http.Request)
	GetSingle(w http.ResponseWriter, r *http.Request)
	PostSingle(w http.ResponseWriter, r *http.Request)
}

func (r *Router) ViewSet(url string, vs ViewSet) {
	// index
	r.add("GET", url, vs.Get)
	r.add("POST", url, vs.Post)

	// singular
	url = url + "/:id"
	r.add("GET", url, vs.GetSingle)
	r.add("POST", url, vs.PostSingle)
	r.add("PUT", url, vs.PutSingle)
	r.add("DELETE", url, vs.DeleteSingle)
}
