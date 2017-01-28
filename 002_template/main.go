package main

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc":       strings.ToUpper,
	"ft":       firstThree,
	"fDateMdy": mod,
}

func init() {
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("templates/tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func mod(t time.Time) string {
	return t.Format("02-01-2006")
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

	data := struct {
		Title  string
		People map[int]person
		Cars   map[int]car
		Date   time.Time
	}{
		"Hello World",
		people,
		cars,
		time.Now(),
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
