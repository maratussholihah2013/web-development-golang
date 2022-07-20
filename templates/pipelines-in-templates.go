package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("pipelinestemplate.html"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func main() {
	//err := tpl.ExecuteTemplate(os.Stdout, "pipelinestemplate.html", time.Now())
	err := tpl.ExecuteTemplate(os.Stdout, "pipelinestemplate.html", time.Now().Format("02-01-2006"))
	if err != nil {
		log.Fatalln(err)
	}
}
