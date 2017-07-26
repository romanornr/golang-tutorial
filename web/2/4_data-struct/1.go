package main

import "text/template"
import "os"
import "log"

var tpl *template.Template

type altcoin struct {
	Name   string
	Symbol string
}

func init() {
	tpl = template.Must(template.ParseFiles("1.gohtml"))
}

func main() {
	Litecoin := altcoin{"Litecoin", "LTC"}

	err := tpl.Execute(os.Stdout, Litecoin)
	if err != nil {
		log.Fatalln(err)
	}
}
