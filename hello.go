package main

import (
	"bufio"
	"fmt"
	"github.com/ApolloTitanGolang/stringutil"
	"io/ioutil" // for the WebReader
	"log"
	"net/http" // for the WebReader
	"os"
	"strconv"
)

// test 11.04.2017 / v2

// variable in package level scope, also in other files available
var e = "package scope"

// constants
const p string = "this is a constant variable"

func main() {
	fmt.Printf("\n-------------------------------------\n")
	fmt.Printf("Hello, world.\n")
	fmt.Printf(stringutil.Reverse("!oG, olleH"))
	fmt.Printf("\n")

	variableTutorial()
	scopeTutorial()
	anonymousFunctionTutorial()
	closerTutorial()
	blankIdentifier()
	memoryAddress()
	callZero()
	remainder()
	loopSample()
	nestedLoopSample()
	loopWithCondition()
	loopForWithNoCondition()
	runeSample()
	switchSample()
	switchOnType("antman")
	switchOnType(42)
	ifSample()
	// userInputSample()
	exercise01()
	functionSample()
	functionVariadicSample()
	functionExpressionSample()
	callbackFunctionSample()
	recursionSample()
	deferSample()
	callByValueSample()
	probExercise1()
	probExercise3()
	probExercise4()
	probExercise5()
	arraySample()
	arraySample2()
	sliceSample()
	mapSample()
	hashTabelSample()

}

/* ------------------------------------------------------------------
	Variables (declare, assign, initialize)
------------------------------------------------------------------*/
func variableTutorial() {

	a := 12
	b := "golang"
	c := 9.81
	d := true

	var e int = 72
	f := 75
	fmt.Println(e)
	fmt.Println(string(e))
	fmt.Println(string(f))

	fmt.Printf("\n\n-variableTutorial--------------------\n")
	fmt.Println(string(a))

	fmt.Printf("%v - %T - %#v - %+v\n", a, a, a, a)
	fmt.Printf("%v - %T - %#v - %+v\n", b, b, b, b)
	fmt.Printf("%v - %T - %#v - %+v\n", c, c, c, c)
	fmt.Printf("%v - %T - %#v - %+v\n", d, d, d, d)
	fmt.Printf("%v", p)
}

/* ------------------------------------------------------------------
	SCOPE (universe, package, file, block)
	- https://golang.org/ref/spec#Declarations_and_scope
	- http://www.golang-book.com/books/web/01-02
------------------------------------------------------------------ */
func scopeTutorial() {
	fmt.Printf("\n-------------------------------------\n")
	fmt.Printf("%v - %T - %#v - %+v", e, e, e, e)

	a := "block scope"
	fmt.Println(a)
	// does not work: fmt.Println(b)
}

/* ------------------------------------------------------------------
	Anonymous Function
------------------------------------------------------------------*/
func anonymousFunctionTutorial() {
	x := 11

	// assign a function to a variable
	increment := func() int {
		x++
		return x
	}
	fmt.Printf("-------------------------------------\n")
	fmt.Println(increment())

}

/* ------------------------------------------------------------------
	Closure
	Closure helps to limit the scope of variables
------------------------------------------------------------------*/
func wrapper() func() int {
	x := 8
	// return an function
	return func() int {
		x++
		return x
	}
}

func closerTutorial() {
	increment := wrapper()
	fmt.Println(increment())
}

/* ------------------------------------------------------------------
	Blank identifier
	Read a html - page
------------------------------------------------------------------*/
func blankIdentifier() {
	// instead of catch errors we throw it with "_" away
	res, _ := http.Get("http://www.google.at")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("\n-------------------------------------\n")
	fmt.Printf("\n\n%s", page)
}

/* ------------------------------------------------------------------
	memory address - pointers
------------------------------------------------------------------*/
func memoryAddress() {
	a := 43

	fmt.Printf("\n\n-------------------------------------")
	fmt.Println("\n\na - ", a)
	// Location, where the value is stored
	fmt.Println("a's memory address - ", &a)
	fmt.Printf("value in dezimal: %d \n", &a)

	// pointers
	var b *int = &a // pointer to the location, b is a int pointer
	fmt.Println(b)  // address of a
	fmt.Println(*b) // value of a - 43 - dereference - give me the value ...

	*b = 42        // value change to 42
	fmt.Println(a) // value a changed
}

/* ------------------------------------------------------------------
	call by reference??
------------------------------------------------------------------*/
func zero(x *int) {
	fmt.Println(x) // x is the memory location
	*x = 0         // dereference to set the value
}

func callZero() {
	x := 5
	fmt.Printf("\n call by ... ---------------------\n")
	fmt.Println(x)
	zero(&x)
	fmt.Println(x)
	fmt.Println(&x)
}

/* ------------------------------------------------------------------
	remainder
------------------------------------------------------------------*/
func remainder() {
	x := 13 / 3
	y := 13 % 3
	z := 10 % 6

	fmt.Printf("\n-------------------------------------\n")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

/* ------------------------------------------------------------------
	CONTROL FLOW
	- sequenze
	- loop / iterative
	- conditionals
------------------------------------------------------------------*/
func loopSample() {
	fmt.Printf("\n-------------------------------------\n")
	for i := 0; i <= 10; i++ {
		fmt.Printf("%v_", i)
	}
}

func nestedLoopSample() {
	fmt.Printf("\n-------------------------------------\n")
	for i := 1; i <= 2; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Println("i:", i, " - j:", j)
		}
		fmt.Println("End of the inner loop")
	}
	fmt.Println("End of the outer loop")
}

func loopWithCondition() {
	i := 0
	fmt.Printf("\n-------------------------------------\n")
	for i < 2 {
		fmt.Println(i)
		i++
	}
}

func loopForWithNoCondition() {
	i := 0
	fmt.Printf("\n-------------------------------------\n")
	for {
		i++
		if i == 2 {
			fmt.Println("2 is reached")
			continue // stop this iteration and continue with the next (3)
		}
		fmt.Println(i)
		if i == 3 {
			fmt.Println("Break the loop")
			break
		}
	}
}

/* ------------------------------------------------------------------
	Type RUNE is a alias for int32
------------------------------------------------------------------*/
func runeSample() {
	fmt.Printf("\nRUNE -------------------------------\n")

	//
	// Show 5 bytes - for each character one byte
	fmt.Println("Hello")
	fmt.Println([]byte("Hello")) // conversion in go / cast in other languages
	fmt.Println([]byte(string(200)))

	b := 72
	c := b + 3
	fmt.Printf("\n int string-------------------------------------\n")
	fmt.Println(rune(b))
	fmt.Println(string(b))
	fmt.Println(string(c))
	fmt.Printf("%T \n", b)
	fmt.Printf("%v", string(10)) // Do a line feed (https://en.wikipedia.org/wiki/ASCII)
	fmt.Println(strconv.Itoa(123451))
	fmt.Println([]byte(strconv.Itoa(123451)))

	fmt.Printf("\n This is a rune-------------------------------------\n")
	foo := 'H' // this is a rune
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	fmt.Printf("%T \n", rune(foo))

	fmt.Printf("\n This is a string-------------------------------------\n")
	fooString := "H"
	fmt.Println(fooString)
	fmt.Printf("%T \n", fooString)

	fmt.Printf("\n ASCII-------------------------------------\n")
	for i := 50; i <= 145; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
}

/* ------------------------------------------------------------------
	switch
------------------------------------------------------------------*/
type contact struct {
	greeting string
	name     string
}

func switchSample() {
	fmt.Printf("\n switch-------------------------------------\n")

	hero := "Superwoman"

	switch hero {
	case "Batman":
		fmt.Println("batman")
	case "Superman", "Superwoman":
		fmt.Println("superman")
		fallthrough
	case "Spiderman":
		fmt.Println("spiderman")
	default:
		fmt.Println("no hero found")
	}

	switch {
	case hero == "Batman":
		fmt.Println("batman")
	case hero == "Superwoman":
		fmt.Println("superwoman")
	default:
		fmt.Println("no hero")
	}
}

func switchOnType(hero interface{}) {
	// asserting - what is the type of hero?
	fmt.Printf("\n switch on type-----------------------------\n")
	switch hero.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}

/* ------------------------------------------------------------------
	if, else if, else
------------------------------------------------------------------*/
func ifSample() {
	fmt.Printf("\n if samples --------------------------------\n")

	b := true

	// initialization statement
	if food := "Chocolate"; b {
		fmt.Println(food)
	}
}

/* ------------------------------------------------------------------
	user input
------------------------------------------------------------------*/
func userInputSample() {
	var name string

	fmt.Printf("\n user input  -------------------------------\n")
	fmt.Print("Please enter your name: ")
	fmt.Scan(&name)

	fmt.Println("Your name is: ", name)

	var numberOne int
	var numberTwo int
	fmt.Print("Please enter a number: ")
	fmt.Scan(&numberOne)
	fmt.Print("Please enter a second number: ")
	fmt.Scan(&numberTwo)
	fmt.Println(numberOne, " / ", numberTwo, " = ", numberOne/numberTwo)

	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			fmt.Print("Jepp", i, " - ")
		}
	}
}

/* ------------------------------------------------------------------
	exercises
------------------------------------------------------------------*/
func exercise01() {
	fmt.Printf("\n\n exercise 01 -------------------------------\n")
	var sum int = 0
	for i := 0; i <= 1000; i++ {
		if i%3 == 0 {
			sum = sum + i
		} else if i%5 == 0 {
			sum = sum + i
		}
	}
	fmt.Println("Sum: ", sum)

	var sum2 int = 0
	for i := 0; i < 1000; i++ {
		switch {
		case i%3 == 0:
			sum2 = sum2 + i
		case i%5 == 0:
			sum2 = sum2 + i
		}
	}
	fmt.Println("Sum2: ", sum2)

	var sum3 int = 0
	for i := 0; i <= 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum3 = sum3 + i
		}
	}

	fmt.Println("Sum3: ", sum3)
}

/* ------------------------------------------------------------------
	functions, parameters, arguments
	greet is declared with parameter
	when calling greet, pass in an argument
------------------------------------------------------------------*/
func functionSample() {
	fmt.Printf("\n\n functions -------------------------------\n")
	greet("Jane", "Doe")
	greet("John", "Doe")

	fmt.Println(greetWithReturn("Tim ", "Doe"))
	fmt.Println(greetWithTwoReturns("James ", "Bond "))
}

func greet(fname, lname string) {
	fmt.Println(fname, lname, "test", "2")
}

func greetWithReturn(fname, lname string) string {
	// Do printing to a string
	return fmt.Sprint(fname, lname)
}

func greetWithTwoReturns(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}

/* ------------------------------------------------------------------
	functions,
	- variadic parameters (...)
	- range (foreach loops)
------------------------------------------------------------------*/
func functionVariadicSample() {

	fmt.Printf("\n\n variadic  -------------------------------\n")
	fmt.Println("Average: ", average(10, 20))

	data := []float64{200, 300}
	n := average(data...) // split the list / slice, so it can be read by average
	fmt.Println("Average: ", n)
}

func average(sf ...float64) float64 {
	total := 0.0
	fmt.Println(sf)

	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}

/* ------------------------------------------------------------------
	functions,
	- func-expression
		- assign a function to a variable
	- closure
------------------------------------------------------------------*/
func functionExpressionSample() {

	fmt.Printf("\n\n function expression -----------------------\n")

	greeting := func() {
		fmt.Println("Helo world!")
	}

	greeting()
	greeting()

	fmt.Printf("\nType: %T \n", greeting)

	// get a function with return-type func() string
	greet := makeGreeterFunction()
	fmt.Println(greet())
	fmt.Printf("%T\n", greet)
}

// return Type is a  "func() string"
func makeGreeterFunction() func() string {
	return func() string {
		return "Hello func() string!"
	}
}

/* ------------------------------------------------------------------
	functions,
	- callback
	https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/4037210
	https://github.com/GoesToEleven/GolangTraining/blob/master/14_functions/12_callbacks/02_filter-nums/main.go
------------------------------------------------------------------*/
func callbackFunctionSample() {
	fmt.Printf("\n\n callback function -----------------------\n")
	visit([]int{1, 2, 3, 4, 5}, func(n int) {
		fmt.Println(n)
	})

	fmt.Printf("\n\n callback function 2----------------------\n")
	visit([]int{1, 2, 3, 4, 5}, func(n int) {
		n = n * n
		fmt.Println(n)
	})
}

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

/* ------------------------------------------------------------------
	recursion
------------------------------------------------------------------*/
func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func recursionSample() {
	fmt.Printf("\n\n recursion sample----------------------\n")
	fmt.Println(factorial(4))
}

/* ------------------------------------------------------------------
	defer
	execute before exit the func
------------------------------------------------------------------*/
func deferSample() {

	fmt.Printf("\n\n defer sample--------------------------\n")
	defer World()
	Hello()
}

func Hello() {
	fmt.Printf("Hello ")
}

func World() {
	fmt.Printf("World")
}

/* ------------------------------------------------------------------
	call by value
	- EVERYTHING is passed by value
		- also pointers (you pass the address - by value)
------------------------------------------------------------------*/
func callByValueSample() {
	fmt.Printf("\n\n call by value sample------------------\n")
	age := 44
	fmt.Println(&age) // get the memory address. 0xc20809d938

	changeMe(&age) // pass the value (which is the address)

	fmt.Println(&age) // 0xc20809d938
	fmt.Println(age)  // 24

	fmt.Printf("\n\n call by value sample (slice)------------\n")
	m := make([]string, 1, 25)
	fmt.Println(m) // []
	changeMySlice(m)
	fmt.Println(m) // [Ben]
}

func changeMe(z *int) { // get the pointer to an int (receive address which is storing an int)
	fmt.Println(z)  // 0xc20809d938
	fmt.Println(*z) // 44 (operator "*" DEREFERENCE the address
	*z = 24
	fmt.Println(z)  // 0xc20809d938
	fmt.Println(*z) // 24
}

// slices, maps and channels are REFERENCE-Types
func changeMySlice(z []string) {
	z[0] = "Ben"
	fmt.Println(z) // [Ben]
}

/* ------------------------------------------------------------------
	Problems
	- Exercise 1,2 (functions)
------------------------------------------------------------------*/
func probExercise1() {
	fmt.Printf("\n\n problem exercise 1------------------\n")

	x := 1
	a, b := halfValue(x)
	fmt.Printf("\nValue: %v - half value: %v", x, a)
	fmt.Printf("\nIs even: %v", b)

	x = 2
	a, b = halfValue2(x)
	fmt.Printf("\nvalue: %v - half value: %v", x, a)
	fmt.Printf("\nIs even: %v", b)

	// func expression version
	x = 4
	halfFunc := func(n int) (float64, bool) {
		return float64(x) / 2, x%2 == 0
	}
	a, b = halfFunc(x)
	fmt.Printf("\nvalue: %v - half value: %v", x, a)
	fmt.Printf("\nIs even: %v", b)

}

func halfValue(x int) (float64, bool) { // long version
	var y float64
	y = float64(x) / 2
	z := false

	if x%2 == 0 {
		z = true
	}

	return y, z
}

func halfValue2(x int) (float64, bool) { // short, better version
	return float64(x) / 2, x%2 == 0
}

/* ------------------------------------------------------------------
	Problems
	- Exercise 3 (variadic parameter)
------------------------------------------------------------------*/
func probExercise3() {
	fmt.Printf("\n\n problem exercise 3------------------\n")
	fmt.Println(findGreatestNumber(1, 2, 7, 90, 5, 91.4, 10)) // 91.4
	fmt.Printf("%T\n", findGreatestNumber())                  // float64
}

func findGreatestNumber(args ...float64) float64 {
	var greatestNumber float64 // slice of float64
	fmt.Printf("%T %v\n", args, args)

	for _, n := range args {
		if n > greatestNumber {
			greatestNumber = n
		}
	}

	return greatestNumber
}

/* ------------------------------------------------------------------
	Problems
	- Exercise 4 (bool expression value)?
------------------------------------------------------------------*/
func probExercise4() {
	fmt.Printf("\n\n problem exercise 4------------------\n")

	// false		false		true
	// true
	if (true && false) || (false && true) || !(false && false) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	// smaller version!
	fmt.Println((true && false) || (false && true) || !(true || true))
}

/* ------------------------------------------------------------------
	Problems
	- Exercise 5 (function - different calls)?
------------------------------------------------------------------*/
func probExercise5() {
	fmt.Printf("\n\n problem exercise 5------------------\n")

	fmt.Println(foo(1, 2))
	fmt.Println(foo(1, 2, 3))

	aSlice := []int{1, 2, 3, 4}
	fmt.Println(foo(aSlice...)) // get the slice values

	fmt.Println(foo())
}

func foo(n ...int) []int {
	return n
}

/* ------------------------------------------------------------------
	Arrays
	- This is an array: 	var x [56]string
	- This is a slice: 	var x []string
------------------------------------------------------------------*/
func arraySample() {
	fmt.Printf("\n\n arrays examples---------------------\n")

	// int
	var y [58]int
	fmt.Println(y) // shows already 58 times 0
	fmt.Println(len(y))
	fmt.Println(y[42])

	// strings
	var x [58]string // shows an empty array
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

}

/* ------------------------------------------------------------------
	Arrays II
------------------------------------------------------------------*/
func arraySample2() {
	fmt.Printf("\n\n arrays examples II---------------------\n")
	var x [256]int

	fmt.Println(len(x))
	fmt.Println(x[42])

	for i := 0; i < 25; i++ {
		x[i] = i * 2
	}

	for i, v := range x {
		fmt.Printf("%v - %v - %T - %b\n", i, v, v, v) // value, type, binary
		// fmt.Printf("%v - %T - %b\n", i, i, i)
		if i > 30 {
			break
		}
	}
}

/* ------------------------------------------------------------------
	Slice
	- https://gobyexample.com/slices
	- https://blog.golang.org/go-slices-usage-and-internals
	- reference type like map
	- slices can change there size
------------------------------------------------------------------*/
func sliceSample() {
	fmt.Printf("\n\n slice examples I---------------------\n")

	// slice 1
	mySlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])    // slicing a slice
	fmt.Println(mySlice[2])      // slice index access
	fmt.Println("myString: "[2]) // string index access

	// slice 2
	// sice of 5, in use 10 (dynamic range up to 5)
	// after 5 it increases the size by a algorithm (double cap, ...)
	mySlice2 := make([]int, 0, 5)
	fmt.Println("---------------------")
	fmt.Println(mySlice2)
	fmt.Println(len(mySlice2)) // what is in the slice: 0
	fmt.Println(cap(mySlice2)) // what is the capacity: 5
	fmt.Println("---------------------")

	// slice 3
	mySlice3 := []string{
		"Good morning",
		"Bonjour",
		"Hallo!",
		"dias",
	}

	for i, entry := range mySlice3 {
		fmt.Println(i, entry)
	}
	fmt.Println("---------------------")

	// slice 4
	mySlice4 := make([]string, 3, 5)
	mySlice4[0] = "Good morning"
	mySlice4 = append(mySlice4, "Hallo")
	mySlice4 = append(mySlice4, "Bonjour")

	for i, entry := range mySlice4 {
		fmt.Println(i, entry)
	}
	fmt.Println("---------------------")

	// 2D slice (dynamic)
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen) // create a dynamic size-slice
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

/* ------------------------------------------------------------------
	Map
	- reference type (pointer)
	- a dictionary
	- https://gobyexample.com/maps
	- http://www.golang-book.com/books/intro/6
------------------------------------------------------------------*/
func mapSample() {
	fmt.Printf("\n\n map examples ---------------------\n")

	m := make(map[string]int) // create a map where key is a string and value is an int
	m["k1"] = 6
	m["k2"] = 12
	m["k3"] = 24
	m["k4"] = 48

	fmt.Println("len: ", len(m))
	delete(m, "k2") // delete the value 12 (key: k2)
	fmt.Println("len: ", len(m))
	fmt.Println("map ", m)

	_, valueThere := m["k2"] // ignore the value itself with a blank identifier
	fmt.Println("is the value here? ", valueThere)

	// var myGreeting = make(map[string]string)	// works
	// myGreeting := map[string]string {}		// works too
	var myGreeting = map[string]string{} // works too
	myGreeting["Tim"] = "God morning"
	myGreeting["Jenny"] = "Bonjour"
	myGreeting["Tom"] = "Bongiorno"
	fmt.Println(myGreeting)
	fmt.Println(myGreeting["Jenny"])

	// delete(myGreeting, "Jenny")

	if val, exists := myGreeting["Jenny"]; exists {
		fmt.Printf("val: %v exists: %v", val, exists)
	} else {
		fmt.Printf("val: %v exists: %v", val, exists)
	}

	// range loop
	for key, val := range myGreeting {
		fmt.Printf("\nkey: %v, val: %v - ", key, val)
	}
}

func hashTabelSample() {
	fmt.Printf("\n\n hashtabel examples ---------------------\n")

	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// Create a map
	words := make(map[string]string)

	// bufio (input-outbut buffer)
	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = ""
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading")
	}

	i := 0
	for k, _ := range words {
		fmt.Println(k)
		if i == 10 {
			break
		}
		i++
	}

}
