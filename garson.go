package garson

import "net/http"

// New creates and return a new router object
// it should be passed to http.ListenAndServe
// example:
// router := garson.New()
// router.Get("/hello", func(w http.ResponseWriter, r *http.Request){})
// http.ListenAndServe(":8080", router)
func New() *Router {
	return &Router{}
}

// GetParam gets the value from route paramters of the key passed
// e.g: if a route was registered like this:
// router.Get("/hello/:name", ...)
// to get the value of (:name) when users access this route:
// name, ok := GetParam(request, "name")
/*func GetParam(r *http.Request, key string) (string, bool) {
	ctx := r.Context()
	params := ctx.Value("route_params").(Params)
	val, ok := params[key]
	return val, ok
}*/

func Handle(method string, path string, handler http.HandlerFunc) {

}
