# Garson

Simple Go router lies on top of net/http package. created for learning purposes.

[![GoDoc](https://godoc.org/github.com/emostafa/garson?status.svg)](https://godoc.org/github.com/emostafa/garson)

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

To get the params, just get the request's context, then check the value
of "route_params" key

```go
func someHandler(w Http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if val := ctx.Value("route_params"); val != nil {
		params := val.(map[string]string)
		fmt.Println(params["id"])
	}
	...
}
```
