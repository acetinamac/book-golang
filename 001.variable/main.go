package main

import (
	"fmt"
	"math"
	"os"
)

var Global int = 1234
var AnotherGlobal int = -5678

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command line argument")
		return
	}

	argument := os.Args[1]

	switch argument {
	case "0":
		fmt.Println("Zero!")
	case "1":
		fmt.Println("One!")
	case "2", "3", "4":
		fmt.Println("2 or 3 or 3")
	default:
		fmt.Println("Other number")
	}

	var j int
	i := Global + AnotherGlobal
	println("Initial j value:", j)
	j = Global

	k := math.Abs(float64(AnotherGlobal))
	fmt.Printf("Global=%d, i=%d j=%d k=%.2f\n", Global, i, j, k)
}
