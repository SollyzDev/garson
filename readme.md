# Garson

Simple Go router created for learning purposes

#### Installation

    go get github.com/eslammostafa/garson

#### Usage

    import (
        ...
        "net/http"
        "github.com/eslammostafa/garson"
    )


    func main() {
        router := garson.New()
        router.Get("/hello", func(w http.ResponseWriter, r *http.Request){
            return "Hello World"
        })
        http.ListenAndServe(":8080", router)
    }

#### TODO

* regexp for url path and extract them as variables in Request
* router.ServeStatic(path string) a function to serve the static files
* write Response helpers
* use my custom response instead of http.ResponseWriter
* Middlewares support !
