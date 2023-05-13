# tiny

## Install

```
go get github.com/julianossilva/tiny 
```

## Usage


```go
package main

import (
    "github.com/julianossilva/tiny"

    "net/http"
    "fmt"
    "log"
)

func main() {

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

	s := &http.Server{
		Addr:           ":8080",
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}

```