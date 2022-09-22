package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
}

//GoDOC package time
func monthDayYear(t time.Time) string {
	//return t.Format("01-02-2006")
	return t.Format(time.Kitchen)
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func main() {

	//err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	t := time.Date(2019, time.May, 13, 00, 00, 1, 1, time.Local)
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", t)
	if err != nil {
		log.Fatalln(err)
	}
}
