package main

import "fmt"

// Storing functions as custom type
type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	moreNumbers := []int{5, 1, 2}

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)

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

// Returning functions from functions
func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
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
