package main

import (
	"io"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Homepage")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "My Sweetie Dog")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Maratus Sholihah")
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
