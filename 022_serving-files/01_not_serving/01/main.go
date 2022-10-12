package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", animal)
	http.ListenAndServe(":8080", nil)
}

func animal(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `
	<!--not serving from our server-->
	<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
	`)
}
