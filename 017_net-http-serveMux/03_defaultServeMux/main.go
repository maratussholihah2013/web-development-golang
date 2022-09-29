package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dogggy dooogggy")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Catttyyyyy")
}

func main() {
	var d hotdog
	var c hotcat

	http.Handle("/dog/", d)
	http.Handle("/cat/", c)
	http.ListenAndServe(":8080", nil)
}
