package garson

import (
	"net/http"
	"testing"
)

func TestBasicRouter(t *testing.T) {
	router := New()
	router.Get("/api/:name", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	_, params, err := router.Try("/api/foo", "GET")
	if err != nil {
		t.Error(err.Error())
	}
	if params["name"] != "foo" {
		t.Error("Invalid value of parameter(:name)")
	}
}

func TestAddingRoutePanicsOnBadInput(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log("Router panics on bad input, instead of returning error")
		}
	}()

	router := New()
	router.Get(`\p{malformed}`, nil)
	router.Try("test", "GET")
}

func TestRoutesAreThreadSafe(t *testing.T) {
	router := New()
	router.Get("/:name", nil)

	_, params1, _ := router.Try("/first", "GET")
	if param := params1["name"]; param != "first" {
		t.Error(`For path "/first", expected :name="first", got`, param)
	}

	_, params2, _ := router.Try("/second", "GET")
	if param := params2["name"]; param != "second" {
		t.Error(`For path "/second", expected :name="second", got`, param)
	}

	if param := params1["name"]; param != "first" {
		t.Error("Thread-unsafe sharing between two routes parameters")
	}
}

// func TestGetParam(t *testing.T) {
// 	router := New()
// 	router.Get("/:name", func(w http.ResponseWriter, r *http.Request) {
// 		name, ok := GetParam(r, "name")
// 		if ok == false {
// 			t.Error("Failed to get parameter (:name)")
// 		}

// 		t.Logf("name is %s", name)

// 		w.Write([]byte("success"))
// 	})

// 	// handler, _, _ := router.Try("/eslam", "GET")
// 	// r, _ := http.NewRequest("GET", "http://localhost:3000/eslam", nil)
// 	// handler(, r)
// 	// httptest.NewServer(router)

// 	t.Log("sending request to router")
// 	r := httptest.NewRequest("GET", "http://localhost:3000/eslam", nil)
// 	handler, _, _ := router.Try("/eslam", "GET")
// 	handler(nil, r)
// }

func TestRouterMiddlewares(t *testing.T) {

}

func TestRouteMiddlewares(t *testing.T) {

}
