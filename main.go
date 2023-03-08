package main

import (
	"fmt"
	_ "fmt"
)

// Number type constraint
type Number interface {
	int64 | float64
}

// Generics in Go; with generics you can declare and use functions or types that are written to work with any of a set of types provided by calling code.

func main() {
	// Initialize two maps of integers and floats
	mapInts := map[string]int64{"first": 34, "second": 12}

	mapFloats := map[string]float64{"first": 12.99, "second": 33.13}

	fmt.Printf("Non-Generic sums : %v and %v\n\n",
		SumInts(mapInts),
		SumFloats(mapFloats))
	fmt.Println("Usage of Generic Function")
	fmt.Printf("Generic sums using integers and floats parameters : %v and %v\n", SumIntsOrFloats(mapInts), SumIntsOrFloats(mapFloats))
	fmt.Printf("Generic sums using type constraint : %v and %v\n", SumNumbers(mapInts), SumNumbers(mapFloats))
}

// SumInts Sum the elements of a map
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the elements of a map
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, val := range m {
		s += val
	}
	return s
}

// SumNumbers adds numbers of map using interface type Number
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats sums the values of map m. It supports both floats and integers
// as map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
