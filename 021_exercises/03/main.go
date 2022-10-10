package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func home(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "tpl.gohtml", "Its Homepage")
	HandleError(res, err)
}

func dog(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "tpl.gohtml", "My sweetie Dog")
	HandleError(res, err)
}

func me(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "tpl.gohtml", "Maratus Sholihah")
	HandleError(res, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {

	http.Handle("/", http.HandlerFunc(home))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}
