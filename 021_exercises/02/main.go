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

	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
