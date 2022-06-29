package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog // underlyng type Handler
	var c hotcat // underlyng type Handler

	mux := http.NewServeMux()
	mux.Handle("/dog/", d) // accepts /dog/*
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}
