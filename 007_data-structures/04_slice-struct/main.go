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
	paud := school{
		Nama:    "Starbright",
		Daerah:  "Mekarwangi",
		Tingkat: "PAUD",
	}

	tk := school{
		Nama:    "Alhafidhii",
		Daerah:  "Soreang",
		Tingkat: "TK A - B",
	}

	sd := school{
		Nama:    "Al Fatih",
		Daerah:  "Katapang",
		Tingkat: "SD",
	}

	smp := school{
		Nama:    "SMP Angkasa",
		Daerah:  "Margahayu",
		Tingkat: "SLTP",
	}

	sma := school{
		Nama:    "SMA Angkasa",
		Daerah:  "Margahayu",
		Tingkat: "SLTA",
	}

	sekolah := []school{paud, tk, sd, smp, sma}

	err := tpl.Execute(os.Stdout, sekolah)
	if err != nil {
		log.Fatalln(err)
	}
}
