package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

// test 22
func main() {
	p1 := Person{"james", "bond", 20}
	p2 := Person{"miss", "moneypenny", 18}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
