package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	triple := createTransformer(3)

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	doubledNumbers := []int{}

	for _, val := range *numbers {
		doubledNumbers = append(doubledNumbers, transform(val))
	}

	return doubledNumbers
}

// Closures: A function that returns a function; the returned function has access to the
// variables of the outer function even after the outer function has finished executing.
// See: https://go.dev/tour/moretypes/25
// This is a closure
func createTransformer(factor int) func(int) int {
	return func(num int) int {
		return num * factor
	}
}
