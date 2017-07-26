package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type altcoin struct {
	Name   string
	Symbol string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles(tpl.gohtml))
}

var fm = template.FuncMap{
	"symbol:": strings.ToUpper,
	"ft":      firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	litecoin := altcoin{
		Name:   "litecoin",
		Symbol: "LTC",
	}

	err := tpl.Execute(os.Stdout, litecoin)
	if err != nil {
		log.Fatalln(err)
	}

}
