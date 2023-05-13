package tiny

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

type Tiny interface {
	http.Handler
}

type Config struct {
	Bag Bag
}

func New(c *Config) Tiny {
	return &tinyImpl{
		Bag: c.Bag,
	}
}

type tinyImpl struct {
	Bag Bag
}

func (t *tinyImpl) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, r := range t.Bag.Routes() {
		if matched, params := match(r, req); matched {

			reqWithParams := requestWithParams(req, params)

			r.Handler().ServeHTTP(w, reqWithParams)
			return
		}
	}

	w.WriteHeader(404)
}

func match(route Route, req *http.Request) (bool, map[string]string) {
	prot := req.Proto
	_ = prot
	if route.Method() != "WEBSOCKET" && route.Method() != req.Method {
		return false, make(map[string]string)
	}

	pathSegments := extractSegments(req.URL.Path)
	patternSegments := extractSegments(route.Pattern())

	if len(pathSegments) != len(patternSegments) {
		return false, make(map[string]string)
	}

	params := make(map[string]string)

	for i, patternSegment := range patternSegments {
		if isParamName(patternSegment) {
			name := strings.Trim(patternSegment, ":")
			params[name] = pathSegments[i]
		} else {
			if patternSegment != pathSegments[i] {
				return false, make(map[string]string)
			}
		}
	}

	return true, params
}

func extractSegments(path string) []string {
	segments := strings.Split(strings.Trim(path, "/"), "/")
	return segments
}

func isParamName(segment string) bool {
	return strings.Index(segment, ":") == 0
}

type TinyContextKey string

func requestWithParams(req *http.Request, params map[string]string) *http.Request {
	key := TinyContextKey("params")

	newContext := context.WithValue(req.Context(), key, params)

	return req.WithContext(newContext)
}

func GetParams(req *http.Request) (map[string]string, error) {
	ctx := req.Context()

	value, ok := ctx.Value(TinyContextKey("params")).(map[string]string)

	if !ok {
		return make(map[string]string), errors.New("params not finded in request context")
	}

	return value, nil
}
