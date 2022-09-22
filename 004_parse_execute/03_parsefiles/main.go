package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	//get data from one.gmao
	tpl, err := template.ParseFiles("one.gmao")
	if err != nil {
		//write the error
		log.Fatalln(err)
	}

	//write standart output of tpl
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln((err))
	}
	//get data from multi files
	tpl, err = tpl.ParseFiles("two.gmao", "zayyan.gmao")
	if err != nil {
		log.Fatalln(err)
	}
	//write standart output from multi parse files
	err = tpl.ExecuteTemplate(os.Stdout, "zayyan.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	//write standart output from multi parse files
	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}
	//execute template bisa juga digunakan
	//std out dr parse
	err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}
	//kalo tanpa reference jadi balik ke yg parse 1
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
