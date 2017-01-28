package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	Name    string
	Country string
}

func main() {

	// string array (composite literal)
	sages := map[int]person{
		0: person{
			"Jesus",
			"Israel",
		},
		1: person{
			"Buddha",
			"India",
		},
	}

	// Create index.html
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	// Write template to index.html
	err = tpl.Execute(nf, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
