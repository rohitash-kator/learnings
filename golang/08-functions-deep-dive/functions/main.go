package main

import "fmt"

// Storing functions as custom type
type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

}

// Functions can be passed as parameters to other functions
// This function takes a slice of integers and a function that transforms each integer and returns a new slice of integers
func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}

	return dNumbers
}

// This is the same as the above function, but it modifies the original slice in place
// func transformNumbers(numbers *[]int, transform func(int) int) {
// 	for index, value := range *numbers {
// 		(*numbers)[index] = transform(value)
// 	}
// }

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}
