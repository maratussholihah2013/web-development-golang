package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/", mountains)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public/pics"))))
	http.ListenAndServe(":8080", nil)
}

func mountains(res http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln("Template didn't execute: ", err)
	}
}
