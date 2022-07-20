package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}
type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	//tpl = template.Must(template.New("parsingfunction.html").Funcs(fm).ParseFiles("./parsingfunction.html"))
	tpl = template.Must(template.ParseFiles("./parsingfunction.html"))
	tpl = tpl.Funcs()
}

func main() {
	g := sage{
		Name:  "Islam",
		Motto: "There is no god but Allah",
	}

	m := sage{
		Name:  "Eleanor Roosevelt",
		Motto: "The future belongs to those who believe in the beauty of their dreams",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{g, m}
	cars := []car{f, c}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
