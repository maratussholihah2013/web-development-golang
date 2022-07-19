package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./parsingslice.html"))
}

func main() {
	schools := []string{"PAUD", "TK A", "TK B", "SD", "SMP", "SMA"}
	err := tpl.Execute(os.Stdout, schools)
	if err != nil {
		log.Fatalln(err)
	}
}
