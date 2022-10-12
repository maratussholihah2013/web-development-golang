package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", mountain)
	http.HandleFunc("/mountain.jpg", picture)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func mountain(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("mountain.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "mountain.gohtml", nil)
}

func picture(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "mountain.jpg")
}
