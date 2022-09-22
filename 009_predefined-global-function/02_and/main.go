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

type user struct {
	Name  string
	Motto string
	Admin bool
}

func main() {

	u1 := user{
		Name:  "Ika",
		Motto: "Maju Tak Gentar",
		Admin: false,
	}

	u2 := user{
		Name:  "",
		Motto: "Merdeka",
		Admin: false,
	}

	u3 := user{
		Name:  "Sholihah",
		Motto: "Hiduplah bangsaku",
		Admin: true,
	}
	users := []user{u1, u2, u3}
	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
