# Garson

Simple Go router lies on top of net/http package. created for learning purposes.

[![GoDoc](https://godoc.org/github.com/emostafa/garson?status.svg)](https://godoc.org/github.com/emostafa/garson)

[![Build Status](https://travis-ci.org/emostafa/garson.svg?branch=master)](https://travis-ci.org/emostafa/garson)

#### Installation

from your shell use "go get" command to install the package

```bash
 go get github.com/emostafa/garson
```

#### Usage

Garson supports 4 http methods, 
first import garson and then initialize the router inside the main func,

```go
import (
    "net/http"
    g "github.com/emostafa/garson"
)


func main() {
    router := g.New()
    router.Get("/posts", func(w http.ResponseWriter, r *http.Request){})
    router.Post("/posts", func(w http.ResponseWriter, r *http.Request){})
    router.Put("/posts/:id", func(w http.ResponseWriter, r *http.Request){})
    router.Delete("/:posts/:id", func(w http.ResponseWriter, r *http.Request){})

    http.ListenAndServe(":8080", router)
}
```

#### Example

```go
import (
    "net/http"
    g "github.com/eslammostafa/garson"
)


func main() {
    router := g.New()

    router.Get("/hello", func(w http.ResponseWriter, r *http.Request){}
        w.Write([]byte("Hello World"))
    })


    http.ListenAndServe(":8080", router)
}
```

#### Route Params

you can easily define params in route by appending a colon ":" then a name,
for example :

```go
router.Get("/api/articles/:id", handler)
```
Garson using go context package to store request parameters.
the requested route parameters are stored with a key named "route_params"

There is a function called "GetParam()" that makes it easier to get those parameters

```go
func someHandler(w Http.ResponseWriter, r *http.Request) {
    id, ok := garson.GetParam(r, "id")
    if ok != false {
		fmt.Println(id)
    }
    ...
}
```

If you prefer to use the context directly, you can access it through the value
of "route_params" key.

```go
func someHandler(w Http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    if val := ctx.Value("route_params"); val != nil {
        params := val.(garson.Params)
        fmt.Println(params["id"])
    }
    ...
}
```


## ViewSets

Garson allows you to create ViewSets and register it to the router directly,
e.g:

Prepare your ViewSet, it's a struct that implements garson.ViewSet interface.
The ViewSet interface itself is built from another interfaces:
 - Indexer
 - Poster
 - Getter
 - Putter
 - Deleter

 You can implement all of them, or some of them, based on your needs.
 Garson will automatically created urls for the methods you have implemented.

```go

type UserViewSet struct {}

func (vs *UserViewSet) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users"))

}
func (vs *UserViewSet) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Just One User"))
}
func (vs *UserViewSet) Post(w http.ResponseWriter, r *http.Request)         {}
```

```go
func main() {
    ...
    vs := &UserViewSet{}
    router.ViewSet("/api/users", vs)
    ....
}
```

The Router will automatically register routes for this ViewSet as following:

```
    GET     /api/users          => vs.Index
    POST    /api/users          => vs.Post
    GET     /api/users/:id      => vs.Get
```
