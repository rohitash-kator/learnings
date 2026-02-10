package main

import "fmt"

func main() {
	age := 32 // Regular variable declaration

	var agePointer *int // Pointer declaration
	agePointer = &age   // Pointer assignment to the address of the age variable

	fmt.Println("Age pointer:", agePointer)
	fmt.Println("Age:", *agePointer) // Dereference the pointer to get the value of the age variable

	// adultYears := getAdultYears(agePointer) // Pass the pointer to the function
	getAdultYears(agePointer) // Pass the pointer to the function
	// fmt.Println("Adult years:", adultYears)
	fmt.Println("Adult years:", age)

}

func getAdultYears(age *int) {
	// return *age - 18 // Dereference the pointer to get the value of the age variable
	*age = *age - 18 // Dereference the pointer to get the value of the age variable
}
