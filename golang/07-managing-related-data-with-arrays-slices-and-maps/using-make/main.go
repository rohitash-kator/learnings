package main

import "fmt"

// Custom type to create a map[string]float64
type floatMap map[string]float64

// Method to output the map
func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	usernames := make([]string, 2, 5)

	// usernames := []string{}

	usernames[0] = "intact.infinity" // This will cause a runtime error of index out of range as the array is empty in case of declared array without using make

	usernames = append(usernames, "rohitash_kator") // Add a new username to the slice
	usernames = append(usernames, "itsrkator")      // Add a new username to the slice

	fmt.Println(usernames) // [rohitash_kator itsrkator]

	fmt.Println(usernames) // [rohitash_kator itsrkator]

	// Using a built-in type map[string]float64
	courseRatings := make(map[string]float64, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.9
	courseRatings["noe"] = 4.6

	fmt.Println(courseRatings)

	// Using a custom type map[string]float64
	ratings := make(floatMap, 3)

	ratings["go"] = 4.7
	ratings["react"] = 4.8
	ratings["angular"] = 4.9
	ratings["noe"] = 4.6

	ratings.output()

	// Using a range to iterate over the slice
	for index, username := range usernames {
		fmt.Println("Index:", index, "Value:", username)
	}

	// Using a range to iterate over the map
	for key, value := range ratings {
		fmt.Println("Key:", key, "Value:", value)
	}
}
