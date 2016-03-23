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

* regexp for urls
