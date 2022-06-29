package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog // underlyng type Handler

	http.Handle("/dog", d)
	http.HandleFunc("/cat", c)
	http.HandleFunc("/bird", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "bird bird bird")
	})

	http.ListenAndServe(":8080", nil) // since is nil use default mux
}
