package main

import "fmt"

func main() {
	result := add(1, 2)

	fmt.Println(result)
}

// Generic function to add two values of any type
func addGeneric[T int | float64 | string](a, b T) T {
	return a + b
}

// Non-generic function to add two values of any type
func add(a, b interface{}) interface{} {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)

	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}

	aString, aIsString := a.(string)
	bString, bIsString := b.(string)

	if aIsString && bIsString {
		return aString + bString
	}

	return nil
}
