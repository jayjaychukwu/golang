package main

import "fmt"

func main() {
	// fmt.Println("Enter a number: ")
	// var input float64
	// fmt.Scanf("%f", &input)
	var input float64
	fmt.Println("Enter a number: ")
	fmt.Scanln(&input)

	output := input * 2

	fmt.Println(output)
}
