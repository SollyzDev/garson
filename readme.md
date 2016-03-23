# GoWeb

a set of tools that i created for purpose of learning

#### Installation

    go get github.com/eslammostafa/goweb

#### Usage

    import (
        ...
        "net/http"
        "github.com/eslammostafa/goweb"
    )


    func main() {
        router := goweb.New()
        router.Get("/hello", func(w http.ResponseWriter, r *http.Request){
            return "Hello World"
        })
    }

#### TODO

* write Response helpers
* use my custom response instead of http.ResponseWriter
* regexp for url path and extract them as variables in Request
* router.ServeStatic(path string) a function to serve the static files
* Middlewares support ! 
