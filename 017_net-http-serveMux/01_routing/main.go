package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "dogggy dooogggy")
	case "/cat":
		io.WriteString(res, "ccaaatttyyy")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
