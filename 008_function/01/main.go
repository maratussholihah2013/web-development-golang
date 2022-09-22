package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type school struct {
	Nama    string
	Daerah  string
	Tingkat string
}

type uniform struct {
	Warna string
	Jenis string
	Level int
}

//create a FuncMap to register functions
//"uc" is what the func will be called in the template
//"uc" is the ToUpper func from package strings
//"ft" slices a string, returning the first three characters
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
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

	satu := uniform{
		Warna: "bebas",
		Jenis: "pendek",
		Level: 1,
	}
	dua := uniform{
		Warna: "merah putih",
		Jenis: "pendek",
		Level: 2,
	}
	tiga := uniform{
		Warna: "biru putih",
		Jenis: "panjang",
		Level: 3,
	}
	empat := uniform{
		Warna: "putih abu-abu",
		Jenis: "panjang",
		Level: 4,
	}

	sekolah := []school{paud, tk, sd, smp, sma}
	seragam := []uniform{satu, dua, tiga, empat}

	//di sini perbedaan dg struct slice struct 01
	data := struct {
		Sekolah []school
		Seragam []uniform
	}{
		Sekolah: sekolah,
		Seragam: seragam,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
