package main

import "text/template"
import "os"
import "log"

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	litecoin := struct {
		High int
		Low  int
	}{
		1,
		2,
	}
	err := tpl.Execute(os.Stdout, litecoin)
	if err != nil {
		log.Fatalln(err)
	}
}
