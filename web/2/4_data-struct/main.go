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
	//altcoins := []string{"Stratis", "Litecoin", "Decred"}

	altcoins := map[string]string{
		"Litecoin": "LTC",
		"Decred":   "DCR",
		"Stratis":  "STRAT"}

	err := tpl.Execute(os.Stdout, altcoins)
	if err != nil {
		log.Fatalln(err)
	}
}
