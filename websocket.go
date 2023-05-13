package tiny

import (
	"net/http"

	"golang.org/x/net/websocket"
)

func WSAdapter(handler func(ws *websocket.Conn)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		websocket.Handler(handler).ServeHTTP(w, r)
	}
}
