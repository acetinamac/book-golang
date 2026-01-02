package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Ordered define los tipos que pueden ser ordenados
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

// QuickSort ordena un slice usando el algoritmo quicksort
func QuickSort[T Ordered](data []T) {
	if len(data) <= 1 {
		return
	}
	quicksortRecursive(data, 0, len(data)-1)
}

// quicksortRecursive implementa la recursión del quicksort
func quicksortRecursive[T Ordered](data []T, low, high int) {
	if low < high {
		// Usar pivote aleatorio para mejor performance en casos extremos
		pivotIndex := partition(data, low, high)
		quicksortRecursive(data, low, pivotIndex-1)
		quicksortRecursive(data, pivotIndex+1, high)
	}
}

// partition reordena el slice y retorna el índice del pivote
func partition[T Ordered](data []T, low, high int) int {
	// Seleccionar pivote aleatorio y moverlo al final
	rand.Seed(time.Now().UnixNano())
	randomIndex := low + rand.Intn(high-low+1)
	data[randomIndex], data[high] = data[high], data[randomIndex]

	pivot := data[high]
	i := low - 1

	for j := low; j < high; j++ {
		if data[j] <= pivot {
			i++
			data[i], data[j] = data[j], data[i]
		}
	}

	data[i+1], data[high] = data[high], data[i+1]
	return i + 1
}

// printSlice imprime un slice con formato
func printSlice[T Ordered](name string, data []T) {
	fmt.Printf("%-15s: %v\n", name, data)
}

func main() {
	// Ejemplo con diferentes tipos de datos
	numbers := []int{33, 10, 55, 26, 19, 87, 41, 2, 68, 75}
	floats := []float64{3.14, 2.71, 1.41, 9.8, 6.28}
	names := []string{"Charlie", "Alice", "Eve", "Bob", "David", "Armin"}

	fmt.Println("=== Antes de ordenar ===")
	printSlice("Números", numbers)
	printSlice("Flotantes", floats)
	printSlice("Nombres", names)

	// Ordenar todos los slices
	QuickSort(numbers)
	QuickSort(floats)
	QuickSort(names)

	fmt.Println("\n=== Después de ordenar ===")
	printSlice("Números", numbers)
	printSlice("Flotantes", floats)
	printSlice("Nombres", names)

	// Demostración con slice vacío y unitario
	empty := []int{}
	single := []int{42}
	QuickSort(empty)
	QuickSort(single)

	fmt.Println("\n=== Casos especiales ===")
	printSlice("Vacío", empty)
	printSlice("Único elemento", single)
}
