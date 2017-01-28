package main

import "fmt"

type secretAgent struct {
	Person
	licenseToKill bool
}

type Person struct {
	Fname string
	lname string
}

type human interface {
	speak()
}

func (p Person) speak() {
	fmt.Println(p.Fname, `says, Good Morning.`)
}

func (s secretAgent) catch() {
	fmt.Println(s.Fname, s.lname, `says, "Shaken, not stirred."`)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	// array of integers
	//xi := []int{2, 4, 7, 9, 42}
	//fmt.Println(xi)
	//
	// Map = Object
	//m := map[string]int{
	//	"Will": 23,
	//	"Beth": 23,
	//}
	//fmt.Println(m)

	// Create person
	p1 := Person{"Will", "Gough"}
	saySomething(p1)
	//p1.speak()

	sa1 := secretAgent{
		Person{
			"James",
			"Bond",
		},
		true,
	}
	saySomething(sa1)
	//sa1.catch()
	//sa1.Person.speak()
}
