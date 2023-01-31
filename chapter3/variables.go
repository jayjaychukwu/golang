package main

import "fmt"

// global scope
var x string = "Hello World"

// constants
const z string = "Hello There"

func main() {
	//local scope
	y := "You!"
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func f() {
	fmt.Println(x)
}

// defining multiple variables
var (
	a = 5
	b = 10
	c = 15
)
