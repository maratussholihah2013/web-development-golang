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

	u1 := struct {
		Score1 int
		Score2 int
	}{5, 9}

	err := tpl.Execute(os.Stdout, u1)
	if err != nil {
		log.Fatalln(err)
	}
}
