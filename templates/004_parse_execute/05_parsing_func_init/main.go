package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	//write standart output of tpl first asc file
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln((err))
	}
	//write standart output from multi parse files
	err = tpl.ExecuteTemplate(os.Stdout, "zayyan.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	//write standart output from multi parse files
	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	//execute template bisa juga digunakan
	//std out dr parse
	err = tpl.ExecuteTemplate(os.Stdout, "uno.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	//kalo tanpa reference jadi balik ke yg parse 1
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
