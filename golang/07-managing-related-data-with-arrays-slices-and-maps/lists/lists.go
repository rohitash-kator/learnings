package main

import "fmt"

func main() {
	fmt.Println("Arrays, Slices and Maps in Go")

	// Array is a fixed length list of values
	// prices := [4]float64{1.01, 2.99, 3.65, 4.67}	// Fixed length array
	var prices [4]float64 = [4]float64{1.01, 2.99, 3.65, 4.67}
	// prices := [...]float64{1.01, 2.99, 3.65, 4.67} // Dynamic length array

	var productNames [5]string = [5]string{"Java: The complete guide."}
	productNames[2] = "Introduction to Data Structures and Algorithms"

	fmt.Println(prices)       // [1.01 2.99 3.65 4.67]
	fmt.Println(productNames) // [Java: The complete guide. Introduction to Data Structures and Algorithms]
	fmt.Println(prices[2])    // 3.65

	// Slice is a dynamic length list of values
	featuredPrices := prices[1:4]         // Slice of prices from index 1 to 3
	featuredPricesFromStart := prices[:4] // Slice of prices from start to index 3
	featuredPricesTillEnd := prices[1:]   // Slice of prices from index 1 to end

	highlightedPrices := featuredPricesFromStart[:1] // Slice from a slice

	fmt.Println(featuredPrices)          // [2.99 3.65 4.67]
	fmt.Println(featuredPricesFromStart) // [1.01 2.99 3.65 4.67]
	fmt.Println(featuredPricesTillEnd)   // [2.99 3.65 4.67]
	fmt.Println(featuredPrices[1])       // 3.65
	fmt.Println(highlightedPrices)       // [1.01]
}
