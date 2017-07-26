package main

import "html/template"
import "os"
import "log"

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "index.tml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
