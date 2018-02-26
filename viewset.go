package garson

import (
	"net/http"
)

// ViewSet interface that represents all the possible methods
// combined together
type ViewSet interface {
	Indexer
	Putter
	Poster
	Getter
	Deleter
}

// Indexer interfaces represents the methods required to handle GET requests in
// CRUD operation to list a resource
type Indexer interface {
	Index(w http.ResponseWriter, r *http.Request)
}

// Putter interfaces represents the methods required to handle PUT requests in CRUD operations
type Putter interface {
	Put(w http.ResponseWriter, r *http.Request)
}

// Poster interfaces represents the methods required to handle POST requests in CRUD operations
type Poster interface {
	Post(w http.ResponseWriter, r *http.Request)
}

// Getter interfaces represents the methods required to handle GET requests in CRUD operations
// to fetch details of a single resource based on id
type Getter interface {
	Get(w http.ResponseWriter, r *http.Request)
}

// Deleter interfaces represents the methods required to handle DELETE requests in CRUD operations
// to delete a single resource based on id
type Deleter interface {
	Delete(w http.ResponseWriter, r *http.Request)
}

// ViewSet creates routes for a given viewset based on the methods that the passed ViewSet
// has implement.
// Example:
//
// type UserViewSet struct {}
// func (vs UserViewSet) Index(w http.ResponseWriter, r *http.Request) {...}
// router.ViewSet("/api/users", &UserViewSet{})
//
// this will create a single route to list users   GET /api/users
func (r *Router) ViewSet(url string, viewset interface{}) {
	if vs, ok := viewset.(Indexer); ok {
		r.add("GET", url, vs.Index)
	}
	if vs, ok := viewset.(Poster); ok {
		r.add("POST", url, vs.Post)
	}
	if vs, ok := viewset.(Getter); ok {
		u := url + "/:id"
		r.add("GET", u, vs.Get)
	}
	if vs, ok := viewset.(Putter); ok {
		u := url + "/:id"
		r.add("PUT", u, vs.Put)
	}
	if vs, ok := viewset.(Deleter); ok {
		u := url + "/:id"
		r.add("DELETE", u, vs.Delete)
	}
}
