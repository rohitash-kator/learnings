package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// 1. Create a new array (!) that contains three hobbies you have
	// Option 1:
	// var hobbies [3]string
	// hobbies = [3]string{"Coding", "Traveling", "Photography"}

	// Option 2:
	hobbies := [3]string{"Coding", "Traveling", "Photography"}

	// Option 3:
	// var hobbies [3]string = [3]string{"Coding", "Traveling", "Photography"}

	// hobbies[0] = "Programming and Coding"

	//	Output (print) that array in the command line.
	fmt.Println(hobbies)

	// 2. Also output more data about that array:
	//	- The first element (standalone)
	fmt.Println(hobbies[0])
	//	- The second and third element combined a new list
	fmt.Println(hobbies[0:2])

	// 3. Create a slice based on the first element that contains
	//	The first and second elements.
	// 	Create that slice in two different ways (i.e. create tow slices in the end)
	// slicedHobbies := hobbies[0:2]
	slicedHobbies := hobbies[:2]
	fmt.Println(slicedHobbies)

	// 4. Re-slice the slice from (3) and change it to contain the second adn last element of the original array.
	reSlicedHobbies := slicedHobbies[1:3]
	fmt.Println(reSlicedHobbies)

	// 5. Create a "dynamic array" that contains your course goals (at least 2 goals)
	goals := []string{"Become Advanced Go Lang developer", "Get promoted to Senior Software Engineer"}
	fmt.Println(goals)

	// 6. Set the second goal to a different one and then add a third goal to that existing dynamic array
	goals[1] = "Get promoted to Software Engineer (I2)"
	goals = append(goals, "Learn and practice for the JavaScript/TypeScript, Node.js/Nest.js, Data Structures and Algorithms, Design Patterns, and System Design interviews and earn new offer letter")
	fmt.Println(goals)

	// 7. Bonus: Create a "Product" struct with title, id, price
	type Product struct {
		Id    int     `json:"id"`
		Title string  `json:"title"`
		Price float64 `json:"price"`
	}
	// 	Create a dynamic list of products (at least 2 products)
	productList := []Product{{Id: rand.Int(), Title: "Go: The complete guide", Price: 45.99}, {Id: rand.Int(), Title: "JavaScript: The complete guide", Price: 99.90}}
	fmt.Println(productList)

	// 	Then add a third product to the existing list of products.
	productList = append(productList, Product{Id: rand.Int(), Title: "Node.js: The complete guide", Price: 79.90})
	fmt.Println(productList)
}
