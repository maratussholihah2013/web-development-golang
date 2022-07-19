package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./main.html"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "main.html", 100)
	if err != nil {
		log.Fatalln(err)
	}
}
