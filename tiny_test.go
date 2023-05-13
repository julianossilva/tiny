package tiny_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julianossilva/tiny"
)

func TestHello(t *testing.T) {

	b := tiny.NewBag()

	b.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	h := tiny.New(&tiny.Config{
		Bag: b,
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "http://example.com/", nil)

	h.ServeHTTP(w, req)

	res := w.Result()
	body, _ := io.ReadAll(res.Body)

	if !strings.Contains(string(body), "Hello World!") {
		t.Fatal("wrong response")
	}

}

func TestParam(t *testing.T) {

	b := tiny.NewBag()

	b.Get("/hello/:name", func(w http.ResponseWriter, r *http.Request) {
		params, err := tiny.GetParams(r)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		name := params["name"]

		w.Write([]byte(fmt.Sprintf("Hello %s!", name)))
	})

	h := tiny.New(&tiny.Config{
		Bag: b,
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "http://example.com/hello/Ana", nil)

	h.ServeHTTP(w, req)

	res := w.Result()
	body, _ := io.ReadAll(res.Body)

	if !strings.Contains(string(body), "Hello Ana!") {
		t.Fatal("wrong response")
	}
}
