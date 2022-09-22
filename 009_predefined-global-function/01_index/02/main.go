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
	xs := []string{"zero", "one", "two", "three", "four"}

	data := struct {
		Number    []string
		ChildName string
	}{
		Number:    xs,
		ChildName: "Zayyan",
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
