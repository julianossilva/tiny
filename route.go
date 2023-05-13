package tiny

import "net/http"

type Route interface {
	Method() string
	Pattern() string
	Handler() http.HandlerFunc
}

func NewRoute(method string, pattern string, handler http.HandlerFunc) Route {
	return &route{
		method:  method,
		pattern: pattern,
		handler: handler,
	}
}

type route struct {
	method  string
	pattern string
	handler http.HandlerFunc
}

func (r *route) Method() string {
	return r.method
}

func (r *route) Pattern() string {
	return r.pattern
}

func (r *route) Handler() http.HandlerFunc {
	return r.handler
}
