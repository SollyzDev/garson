# Garson

Simple Go router created for learning purposes

#### Installation

from your shell use "go get" command to install the package

```bash
 go get github.com/eslammostafa/garson
```

#### Usage

Garson supports 4 http methods, 
first import garson and then initialize the router inside the main func,

```go
import (
    "net/http"
    g "github.com/eslammostafa/garson"
)


func main() {
    router := g.New()
    router.Get("/posts", func(ctx *g.Context){})
    router.Post("/posts", func(ctx *g.Context){})
    router.Put("/posts/:id", func(ctx *g.Context){})
    router.Delete("/:posts/:id", func(ctx *g.Context{})

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
    router.Get("/hello", func(ctx *g.Context) {
        ctx.Write("Hello World")
    })

    // or better send a json response easily
    router.Get("/json", func(ctx *g.Context) {
        // prepare the body you want to be sent as json
        body := map[string]interface{}
        body["item1"] = "this is a nice item"
        body["item2"] = "this is just a nicer item"
        // send a json response using ctx.Json()
        ctx.Json(body)
    })

    http.ListenAndServe(":8080", router)
}
```

#### Context

Garson's context is just a wrapper for the net/http request and response writer.
Context addes useful methods around the default responsewriter, as you 
can easily send Json response for example.

current available methods

```go
func some_handler(ctx *g.Context) {
   // return 404 error
   ctx.NotFound()
   // return string
   ctx.Write("Hello, World!")
   // return json, pass an object and it will be converted to json
   posts := []string{"post1", "post2", "post3"}
   body = map[string]interface{}
   body["posts"] = posts
   ctx.Json(body)
}
```


#### Route Params

you can easily define params in route by appending a colon ":" then a name,
for example :

```go
router.Get("/api/articles/:id", handler)
```

now in your context in the RouteParams you will find a key called "id"
with a string value.

```go
id := ctx.RouteParams["id"]
```

but there is a nicer way to get params using a method called **GetParam**, and supply
it with a default value in case the key was not found.

```go
id := ctx.GetParam("id, 0)
``
