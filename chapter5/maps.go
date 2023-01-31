package main

import "fmt"

func main() {
	//declaration style 1
	//var x map[string]int

	//maps need to be initialized before they can be used
	x := make(map[string]int)
	x["key"] = 10
	x["hello"] = 20
	fmt.Println(x["key"])
	delete(x, "key")
	fmt.Println(x)

	//shorter way to create maps
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon"}

	name, ok := elements["Un"]
	fmt.Println(name, ok)
}
