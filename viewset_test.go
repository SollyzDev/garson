package garson

import (
	"net/http"
	"testing"
)

type IndexerViewSet struct{}

func (vs IndexerViewSet) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index"))
}

func TestIndexerViewSet(t *testing.T) {
	r := New()
	r.ViewSet("/testing/", &IndexerViewSet{})

	h, _, err := r.Try("/testing/", "GET")
	if err != nil {
		t.Error(err)
	}
	t.Log(h)
}

type IndexerGetterViewSet struct{}

func (vs IndexerGetterViewSet) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index"))
}
func (vs IndexerGetterViewSet) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("single get"))
}

func TestIndexerGetterViewSet(t *testing.T) {
	r := New()
	r.ViewSet("/testing", &IndexerGetterViewSet{})

	h, _, err := r.Try("/testing/1", "GET")
	if err != nil {
		t.Error(err)
	}
	t.Log(h)
}

type IndexerPosterViewSet struct{}

func (vs IndexerPosterViewSet) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index"))
}
func (vs IndexerPosterViewSet) Post(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("post"))
}

func TestIndexerPosterViewSet(t *testing.T) {
	r := New()
	r.ViewSet("/testing", &IndexerPosterViewSet{})

	h, _, err := r.Try("/testing", "POST")
	if err != nil {
		t.Error(err)
	}
	t.Log(h)
}

type DeleterViewSet struct{}

func (vs DeleterViewSet) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleted"))
}

func TestDeleterViewSet(t *testing.T) {
	r := New()
	r.ViewSet("/testing", &DeleterViewSet{})

	h, _, err := r.Try("/testing/1", "DELETE")
	if err != nil {
		t.Error(err)
	}
	t.Log(h)
}
