package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type school struct {
	Nama    string
	Daerah  string
	Tingkat string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	sekolah := school{
		Nama:    "Starbright",
		Daerah:  "Mekarwangi",
		Tingkat: "PAUD",
	}
	err := tpl.Execute(os.Stdout, sekolah)
	if err != nil {
		log.Fatalln(err)
	}
}
