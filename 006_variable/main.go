package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	//parsing data
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Learning Golang for web development")
	if err != nil {
		log.Fatalln(err)
	}
}
