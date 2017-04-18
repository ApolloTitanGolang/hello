package main

import "fmt"


func main() {
	// input dataset
	x := [4][3]int {
		{0,0,1},
		{0,1,1},
		{1,0,1},
		{1,1,1}}

	// output dataset
	y := [4]int {0,1,1,0}

	fmt.Println(x)
	fmt.Println(y)
}