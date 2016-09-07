# Garson

Simple Go router created for learning purposes

#### Installation

from your shell use "go get" command to install the package

```bash
 go get github.com/eslammostafa/garson
```

#### Usage

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

#### TODO
* write documentation for every struct and method
* regexp for url path and extract them as variables in Request
* router.ServeStatic(path string) a function to serve the static files
* write Response helpers
* use my custom response instead of http.ResponseWriter
* Middlewares support !
