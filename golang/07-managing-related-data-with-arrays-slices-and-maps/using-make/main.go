package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	usernames := make([]string, 2, 5)

	// usernames := []string{}
	// usernames[0] = "rohitash-kator" // This will cause a runtime error of index out of range as the array is empty in case of declared array without using make

	usernames = append(usernames, "rohitash_kator") // Add a new username to the slice
	usernames = append(usernames, "itsrkator")      // Add a new username to the slice

	fmt.Println(usernames) // [rohitash_kator itsrkator]

	fmt.Println(usernames) // [rohitash_kator itsrkator]

	courseRatings := make(map[string]float64, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.9
	courseRatings["noe"] = 4.6

	fmt.Println(courseRatings)

	ratings := make(floatMap, 3)
	ratings["go"] = 4.7
	ratings["react"] = 4.8
	ratings["angular"] = 4.9
	ratings["noe"] = 4.6

	ratings.output()

}
