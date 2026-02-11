package main

import "fmt"

func main() {
	fmt.Println("Arrays, Slices and Maps in Go")

	prices := [4]float64{1.01, 2.99, 3.65, 4.67}

	var productNames [5]string = [5]string{"Java: The complete guide."}
	productNames[2] = "Introduction to Data Structures and Algorithms"

	fmt.Println(prices)
	fmt.Println(productNames)
	fmt.Println(prices[2])
}
