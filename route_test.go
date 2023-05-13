package tiny_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/julianossilva/tiny"
)

func TestAccessors(t *testing.T) {
	h := func(w http.ResponseWriter, req *http.Request) {}

	r := tiny.NewRoute("GET", "/", h)

	if r.Method() != "GET" {
		t.Fatal("wrong method")
	}

	if r.Pattern() != "/" {
		t.Fatal("worng pather")
	}

	if reflect.ValueOf(r.Handler()).Pointer() != reflect.ValueOf(h).Pointer() {
		t.Fatal("wrong handler")
	}
}
