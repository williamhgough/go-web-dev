package main

import (
	"html/template"
	"log"
	"os"
	"strings"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc":       strings.ToUpper,
	"ft":       firstThree,
	"fDateMdy": mod,
}

func init() {
	tpl = template.Must(template.New("index.gohtml").Funcs(fm).ParseGlob("templates/*.gohtml"))
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

type page struct {
	Title   string
	Heading string
	Input   string
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
		People map[int]person
		Cars   map[int]car
		Date   time.Time
		Page   page
	}{
		people,
		cars,
		time.Now(),
		page{
			"Nothing Escaped",
			"Nothing is escaped with text/template",
			`<script>alert("Yo!")</script>`,
		},
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
