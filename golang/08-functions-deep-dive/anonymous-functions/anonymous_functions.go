package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	// Anonymous function: A function without a name
	transformed := transformNumbers(&numbers, func(num int) int {
		return num * 2
	})

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	doubledNumbers := []int{}

	for _, val := range *numbers {
		doubledNumbers = append(doubledNumbers, transform(val))
	}

	return doubledNumbers
}
