package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q") // query localhost:8080?q=param
	fmt.Fprintln(w, "Do my search: "+v)
}

/* URL Structure
https://video.google.com:80/videoplay?id=788279871928379&lang=en#button

protocol://subdomain.domain:port/path?query=parameter#fragment
*/
