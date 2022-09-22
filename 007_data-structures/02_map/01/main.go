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
	schools := map[string]string{
		"dini":      "PAUD",
		"taman1":    "TK A",
		"taman2":    "TK B",
		"dasar":     "SD",
		"menengah":  "SMP",
		"menengah2": "SMA",
	}
	err := tpl.Execute(os.Stdout, schools)
	if err != nil {
		log.Fatalln(err)
	}
}
