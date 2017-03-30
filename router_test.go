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

// func TestAddingRoutePanicsOnBadInput(t *testing.T) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Error("Router panics on bad input, instead of returning error")
// 		}
// 	}()
//
// 	router := New()
// 	router.Get(`\p{malformed}`, nil)
// 	router.Try("test", "GET")
// }

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
