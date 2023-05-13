package tiny_test

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julianossilva/tiny"
	"golang.org/x/net/websocket"
)

func TestHelloWS(t *testing.T) {

	b := tiny.NewBag()

	b.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	b.WS("/ws", tiny.WSAdapter(func(c *websocket.Conn) {
		io.Copy(c, c)
	}))

	h := tiny.New(&tiny.Config{
		Bag: b,
	})

	/**
	* HTTP
	 */

	ts := httptest.NewServer(h)
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if !strings.Contains(string(greeting), "Hello World!") {
		t.Fatal("wrong response")
	}

	/**
	* DIAL
	 */
	origin := ts.URL // "http://localhost/"
	url := strings.Replace(fmt.Sprintf("%s/ws", origin), "http://", "ws://", 1)
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte("Hello Websocket!\n")); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}

	msgString := string(msg[:n])
	if !strings.Contains(msgString, "Hello Websocket!") {
		t.Fatal("wrong message")
	}
}
