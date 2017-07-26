package main

import "html/template"
import "os"
import "log"

type course struct {
	Number string
	Name   string
	Units  int
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Introduction to golang", 4},
				course{"CSCI-130", "Introduction to C++", 3},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"CSCI-40", "Introduction to golang", 4},
				course{"CSCI-130", "Introduction to C++", 3},
			},
		},
	}
	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}
