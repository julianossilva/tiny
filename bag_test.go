package tiny_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/julianossilva/tiny"
)

func TestCreateANewBag(t *testing.T) {
	h1 := func(w http.ResponseWriter, req *http.Request) {}
	h2 := func(w http.ResponseWriter, req *http.Request) {}

	b := tiny.NewBag()

	if b.Size() != 0 {
		t.Fatal("bag size error")
	}

	b.Get("/", h1)

	if b.Size() != 1 {
		t.Fatal("bag size error")
	}

	b.Post("/resource", h2)

	if b.Size() != 2 {
		t.Fatal("bag size error")
	}

	rs := b.Routes()

	if rs[0].Method() != "GET" {
		t.Fatal("method error")
	}
	if rs[0].Pattern() != "/" {
		t.Fatal("pattern error")
	}
	if reflect.ValueOf(rs[0].Handler()).Pointer() != reflect.ValueOf(h1).Pointer() {
		t.Fatal("handler error")
	}

	if rs[1].Method() != "POST" {
		t.Fatal("method error")
	}
	if rs[1].Pattern() != "/resource" {
		t.Fatal("pattern error")
	}
	if reflect.ValueOf(rs[1].Handler()).Pointer() != reflect.ValueOf(h2).Pointer() {
		t.Fatal("handler error")
	}
}
