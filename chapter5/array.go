package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	random()
	getIt()
	arrays()
}

func random() {
	var x [5]float64
	x[0] = 91
	x[1] = 89
	x[2] = 21
	x[3] = 83
	x[4] = 60.3

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}

// special type of for loop
func getIt() {
	var x [5]float64
	x[0] = 91
	x[1] = 89
	x[2] = 21
	x[3] = 83
	x[4] = 60.3

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}

// shorter syntax for creating arrays
func arrays() {
	x := [5]float64{
		78,
		90,
		31,
		33,
		44,
	}

	fmt.Println(x[3])
}
