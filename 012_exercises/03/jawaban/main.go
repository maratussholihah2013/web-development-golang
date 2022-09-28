package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Nama, Description string
	Harga             int
}

type menu struct {
	Jenis string
	Items []item
}

type daftarMenu []menu

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	data := daftarMenu{
		menu{
			Jenis: "Sarapan",
			Items: []item{
				item{"Bubur Ayam", "Nasi yang dilumatkan dengan taburan ayam", 10000},
				item{"Teh Hangat", "Segelas teh dengan gula", 5000},
				item{"Nasi Kuning", "Nasi berwarna kuning dengan berbagai pilihan toping", 12000},
				item{"Nasi Uduk", "Nasi gurih yang mengenyangkan", 10000},
			},
		},
		menu{
			Jenis: "Makan Siang",
			Items: []item{
				item{"Gado-gado", "Potongan lontong dengan sayur dan tahu yang disiram kuah", 10000},
				item{"Ayam Penyet", "Ayam goreng dengan lalapan dan sambal", 20000},
				item{"Ikan bakar", "Ikan nila yang masih segar", 15000},
				item{"Es Jeruk", "Segelas air dan es yang diberi perasan jeruk manis", 10000},
			},
		},
		menu{
			Jenis: "Makan Malam",
			Items: []item{
				item{"Kwetiau", "Mie goreng lengkap dengan potongan ayam dan telur", 10000},
				item{"Ayam Penyet", "Ayam goreng dengan lalapan dan sambal", 5000},
				item{"Ikan bakar", "Ikan nila yang masih segar", 15000},
				item{"Es Jeruk", "Segelas air dan es yang diberi perasan jeruk manis", 10000},
			},
		},
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

}
