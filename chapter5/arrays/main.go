package main

import "fmt"

func main() {
	var x [5]int //Creates an array of 5 integers
	x[4] = 100   // assing a value of 100 in array position 4 [array starts from 0]
	//fmt.Println(x) // [0 0 0 0 100]
	fmt.Println(x[4]) // Print the item in the array position of 4
}
