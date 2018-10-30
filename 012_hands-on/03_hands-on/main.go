package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("hotels.html"))
}

type Hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

func main() {
	caliHotels := []Hotel{
		Hotel{"Sunshine Hotel", "37 Sunshine Rd", "LA", 11111, "Southern"},
		Hotel{"Moonlight Hotel", "98 Moonlight Rd", "LA", 11111, "Southern"},
		Hotel{"Starlight Hotel", "24 Starlight Rd", "SF", 22222, "Central"},
	}
	err := tpl.Execute(os.Stdout, caliHotels)
	if err != nil {
		log.Fatalln(err.Error())
	}

}
