package main

import "fmt"

//Printing Average of a slice of numbers
func main() {
	xs := []float64{98, 93, 77, 82, 83}
	total := 0.0
	for _, v := range xs {
		total += v
	}
	fmt.Println(total / float64(len(xs)))
}
