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

type car struct {
	Name         string
	Manufacturer string
	Age          int
}

type items struct {
	People map[int]person
	Cars map[int]car
}

func main() {

	// string array (composite literal)
	people := map[int]person{
		0: person{
			"Jesus",
			"Israel",
		},
		1: person{
			"Buddha",
			"India",
		},
	}

	// Create map of cars
	cars := map[int]car{
		0: car{
			"Mustang",
			"Ford",
			25,
		},
		1: car{
			"Freelander",
			"Land Rover",
			20,
		},
	}

	data := items{
		people,
		cars,
	}

	// Create index.html
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	// Write template to index.html
	err = tpl.Execute(nf, data)
	if err != nil {
		log.Fatalln(err)
	}
}
