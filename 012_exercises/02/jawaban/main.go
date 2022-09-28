package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip, Region string
}

type region struct {
	Region string
	Hotels []hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	data := []region{
		region{
			Region: "Southern",
			Hotels: []hotel{
				hotel{
					Name:    "Kopo Ceria",
					Address: "Jalan Kopo Bihbul 12",
					City:    "Bandung",
					Zip:     "52736",
					Region:  "Southern",
				},
				hotel{
					Name:    "Kopo Beta",
					Address: "Jalan Kopo Katapang 10",
					City:    "Bandung",
					Zip:     "52736",
					Region:  "Southern",
				},
			},
		},
		region{
			Region: "Central",
			Hotels: []hotel{
				hotel{
					Name:    "Kopo Alfa",
					Address: "Jalan Kopo Bihbul 12",
					City:    "Bandung",
					Zip:     "52736",
					Region:  "Southern",
				},
				hotel{
					Name:    "Kopo Giga",
					Address: "Jalan Kopo Katapang 10",
					City:    "Bandung",
					Zip:     "52736",
					Region:  "Southern",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
