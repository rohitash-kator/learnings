package main

import "fmt"

func dynamicListsWithSlices() {
	// Building dynamic Lists with Slices
	prices := []float64{10.99, 8.99}

	fmt.Println(prices[:1])

	prices[1] = 9.99
	// prices[2] = 11.99 // This will cause a runtime error of index out of range as the array is fixed length

	prices = append(prices, 5.99) // This will create a new array with the new price and return the new slice

	fmt.Println(prices) // [10.99 9.99 5.99] [10.99 9.99]
}

func main() {
	fmt.Println("Arrays, Slices and Maps in Go")

	dynamicListsWithSlices()

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
	// Slice is a view of the array, so if you change the value of the slice, it will change the value of the array
	// Slice is a reference to the array, so if you change the value of the array, it will change the value of the slice
	// Slice is a reference to the array, so if you change the value of the array, it will change the value of the slice
	featuredPrices := prices[1:4]         // Slice of prices from index 1 to 3
	featuredPricesFromStart := prices[:4] // Slice of prices from start to index 3
	featuredPricesTillEnd := prices[1:]   // Slice of prices from index 1 to end

	highlightedPrices := featuredPricesFromStart[:1] // Slice from a slice

	fmt.Println(featuredPrices)          // [2.99 3.65 4.67]
	fmt.Println(featuredPricesFromStart) // [1.01 2.99 3.65 4.67]
	fmt.Println(featuredPricesTillEnd)   // [2.99 3.65 4.67]
	fmt.Println(featuredPrices[1])       // 3.65
	fmt.Println(highlightedPrices)       // [1.01]

	featuredPrices[0] = 99.999999 // Change the value of the first element of the featuredPrices slice and the prices array

	fmt.Println(prices) // [1.01 99.999999 3.65 4.67]

	fmt.Println(len(featuredPrices), cap(featuredPrices))
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
}
