package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

// Passing complex struct to template
// Using functions in templates (FuncMap needs to be created)

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"aw": attachWord,
}

type course struct {
	Number string
	Name   string
	Units  string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	AcaYear string
	Fall    semester
	Spring  semester
	Summer  semester
}

func attachWord(word string, s string) string {
	return s + " " + word
}

var tpl *template.Template

// init sets up the templates before main runs
// funcs need to be passed before template files are parsed
// otherwise error is thrown
func init() {
	// name must be set when using Execute, can be empty when using ExecuteTemplate, because there naem must be given
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("tpl/*.gohtml"))
}

func main() {
	years := []year{
		year{
			AcaYear: "2020-2021",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					course{"CSCI-40", "Introduction to Programming in Go", "4"},
					course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					course{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: semester{
				Term: "Spring",
				Courses: []course{
					course{"CSCI-50", "Advanced Go", "5"},
					course{"CSCI-190", "Advanced Web Programming with Go", "5"},
					course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
		},
		year{
			AcaYear: "2021-2022",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					course{"CSCI-40", "Introduction to Programming in Go", "4"},
					course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					course{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: semester{
				Term: "Spring",
				Courses: []course{
					course{"CSCI-50", "Advanced Go", "5"},
					course{"CSCI-190", "Advanced Web Programming with Go", "5"},
					course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", years)
	if err != nil {
		log.Fatalln(err)
	}
}
