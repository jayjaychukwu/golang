package main

import "fmt"

func main() {
	//An array with no determined length
	//var x []float64

	//This is another way to declare a slice
	//x := make([]float64, 5)

	//Another way to declare a slice is this
	//x := make([]float64, 5, 10)

	arr := [5]float64{1, 2, 3, 4, 5}
	x := arr[0:4]
	fmt.Println(x)
}
