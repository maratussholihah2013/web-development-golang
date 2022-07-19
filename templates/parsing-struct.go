package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type paud struct {
	Nama    string
	Daerah  string
	Tingkat string
}

func init() {
	tpl = template.Must(template.ParseFiles("./parsingstruct.html"))
}

func main() {
	sekolah := paud{
		Nama:    "Starbright",
		Daerah:  "Mekarwangi",
		Tingkat: "PAUD",
	}
	err := tpl.Execute(os.Stdout, sekolah)
	if err != nil {
		log.Fatalln(err)
	}
}
