package tiny

import "net/http"

type Bag interface {
	Size() int
	Add(bag Bag)
	Routes() []Route

	Method(method string, pattern string, handler http.HandlerFunc)
	Get(pattern string, handler http.HandlerFunc)
	Head(pattern string, handler http.HandlerFunc)
	Post(pattern string, handler http.HandlerFunc)
	Put(pattern string, handler http.HandlerFunc)
	Delete(pattern string, handler http.HandlerFunc)
	Patch(pattern string, handler http.HandlerFunc)
	Connect(pattern string, handler http.HandlerFunc)
	Options(pattern string, handler http.HandlerFunc)
	Trace(pattern string, handler http.HandlerFunc)

	WS(pattern string, handler http.HandlerFunc)
}

func NewBag() Bag {
	return &bag{
		routes: make([]Route, 0),
	}
}

type bag struct {
	routes []Route
}

func (b *bag) Method(method string, pattern string, handler http.HandlerFunc) {
	r := NewRoute(method, pattern, handler)

	b.routes = append(b.routes, r)
}

func (b *bag) WS(pattern string, handler http.HandlerFunc) {
	b.Method("WEBSOCKET", pattern, handler)
}

func (b *bag) Get(pattern string, handler http.HandlerFunc) {
	b.Method("GET", pattern, handler)
}

func (b *bag) Head(pattern string, handler http.HandlerFunc) {
	b.Method("HEAD", pattern, handler)
}

func (b *bag) Post(pattern string, handler http.HandlerFunc) {
	b.Method("POST", pattern, handler)
}

func (b *bag) Put(pattern string, handler http.HandlerFunc) {
	b.Method("PUT", pattern, handler)
}

func (b *bag) Delete(pattern string, handler http.HandlerFunc) {
	b.Method("DELETE", pattern, handler)
}

func (b *bag) Patch(pattern string, handler http.HandlerFunc) {
	b.Method("PATCH", pattern, handler)
}

func (b *bag) Connect(pattern string, handler http.HandlerFunc) {
	b.Method("CONNECT", pattern, handler)
}

func (b *bag) Options(pattern string, handler http.HandlerFunc) {
	b.Method("OPTIONS", pattern, handler)
}

func (b *bag) Trace(pattern string, handler http.HandlerFunc) {}

func (b *bag) Size() int {
	return len(b.routes)
}

func (b *bag) Add(other Bag) {
	b.routes = append(b.routes, other.Routes()...)
}

func (b *bag) Routes() []Route {
	return b.routes
}

func JoinBags(bags ...Bag) Bag {
	newBag := NewBag()

	for _, b := range bags {
		newBag.Add(b)
	}

	return newBag
}
