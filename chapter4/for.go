package main

import "fmt"

func main() {
	//This is one way to write it
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}

	//This is another way to write it
	for g := 1; g <= 10; g++ {
		fmt.Println(g)
	}
}
