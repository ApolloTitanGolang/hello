package main

import "fmt"

type Person struct {
	first string
	last string
	age int
}

// Embedding an object
type DoubleZero struct {
	Person
	first string
	LicenseToKill bool
}

func main() {
	//
	p1 := DoubleZero{
		Person: Person{
			first: "James",
			last: "Bond",
			age: 20,
		},
		first: "Reborn",	// overwriting parameter <first>
		LicenseToKill: true,
	}

	p2 := Person{"miss", "moneypenny", 18}

	fmt.Println(p1, p1.first, p1.last, p1.Person.first)
	fmt.Println(p2.first, p2.last, p2.age)
}