package main

import (
	"fmt"
)

func main() {
	x := make(map[string]int)
	// var x map[string]int // it will throw a runtime error, map has to be initialized
	x["key"] = 10
	fmt.Println(x["key"])
}
