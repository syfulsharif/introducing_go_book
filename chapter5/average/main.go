package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i := 0; i < len(x); i++ { // len(x) for length of the array x
		total += x[i]
	}
	fmt.Println(total / float64(len(x))) // float64() converts the integer  output of len(x)
}
