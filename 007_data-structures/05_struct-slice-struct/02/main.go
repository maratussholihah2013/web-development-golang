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

type uniform struct {
	Warna string
	Jenis string
	Level int
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
