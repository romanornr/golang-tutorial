package main

import (
	"html/template"
	"log"
	"net/http"
)

type altcoin int

func (a altcoin) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.html", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	var a altcoin
	http.ListenAndServe(":8080", a)
}
