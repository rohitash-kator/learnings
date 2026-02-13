package main

import "fmt"

func main() {
	fact := factorial(5)

	fmt.Println(fact)
}

// Recursion: A function that calls itself (Recursion is a approach to solve problems by breaking them down into smaller sub-problems)
// factorial of a number is the product of all positive integers less than or equal to that number
// 5! = 5 * 4 * 3 * 2 * 1 = 120
// 0! = 1 => Base case (n! = n * (n-1)! for n > 0)[1! = 1 * 0! =>  0! = 1]
// 1! = 1
func factorial(num int) int {
	// Iterative approach
	// result := 1
	// for i := 1; i <= num; i++ {
	// 	result *= i
	// }
	// return result

	// Recursive approach
	if num == 1 || num == 0 {
		return 1
	}

	return num * factorial(num-1)
}
