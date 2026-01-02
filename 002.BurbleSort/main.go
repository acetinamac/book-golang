package main

import "fmt"

type Ordered interface {
	~float64 | ~int | ~string
}

func bubbleSort[T Ordered](data []T) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func main() {
	// Example usage
	ints := []int{64, 34, 25, 12, 22, 11, 90}
	bubbleSort(ints)
	fmt.Println("Sorted integers:", ints)

	names := []string{"John", "Alice", "Bob", "Diana", "Armin"}
	bubbleSort[string](names)
	fmt.Println(names)
}
