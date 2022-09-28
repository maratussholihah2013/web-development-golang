package main

import (
	"fmt"
	"net/http"
)

type zay int

func (z zay) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Apapun code-nya")
}

func main() {
	var z zay
	http.ListenAndServe(":8080", z)
}
