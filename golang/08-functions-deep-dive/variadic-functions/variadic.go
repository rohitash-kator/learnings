package main

import "fmt"

func main() {
	numArr := []int{121, 24, 32, 49}

	// Passing a slice to a variadic function: use ...(SPREAD OPERATOR) after the slice to split it into individual arguments.
	// Same as JavaScript spread at call site: sumUp(10, ...numArr)
	sum := sumUp(10, numArr...)
	fmt.Println(sum)

	anotherSum := sumUp(34, 53, 20, 45)
	fmt.Println(anotherSum)
}

// Variadic function: A function that can be called with any number of trailing arguments.
// The final parameter has type prefixed with ...; inside the function it has type []T (slice of type T where T is the type of the arguments, i.e. Generic type, int, float64, string, etc.).
// Similar to JavaScript:
//   - In the signature: numbers ...int (COLLECTOR OPERATOR) is like rest params (...numbers) — collects trailing args into a slice/array.
//   - At the call site: slice...(SPREAD OPERATOR) unpacks a slice into arguments, like the spread operator (...arr) in JS.
func sumUp(startingNumber int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return startingNumber + sum
}
