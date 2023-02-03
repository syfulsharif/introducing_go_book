package main

import "fmt"

func main() {
	var x [5]float64 //shorthand for array var x := [5]float64{98, 93, 77, 82, 83,}
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	var total float64 = 0
	for _, value := range x { // single "_" is used to tell the compiler we are using temporary variable
		total += value
	}
	fmt.Println(total / float64(len(x)))
}
