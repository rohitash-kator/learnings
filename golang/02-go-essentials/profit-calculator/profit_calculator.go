package main

import "fmt"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// fmt.Print("Enter revenue: ")
	// fmt.Scan(&revenue)

	// fmt.Print("Enter expenses: ")
	// fmt.Scan(&expenses)

	// fmt.Print("Enter tax rate: ")
	// fmt.Scan(&taxRate)

	revenue := getUserInput("Enter revenue: ")
	expenses := getUserInput("Enter expenses: ")
	taxRate := getUserInput("Enter tax rate: ")

	// earningBeforeTax := revenue - expenses
	// earningAfterTax := earningBeforeTax * (1 - taxRate/100)
	// ratio := earningBeforeTax / earningAfterTax

	earningBeforeTax, earningAfterTax, ratio := calculateFinancial(revenue, expenses, taxRate)

	// fmt.Print("Earnings before tax: ", earningBeforeTax)
	// fmt.Print("\nEarnings after tax (Profit): ", earningAfterTax)
	// fmt.Print("\nEBT/profit ratio: ", ratio)

	fmt.Printf("Earnings before tax: %.1f\n", earningBeforeTax)
	fmt.Printf("Earnings after tax (Profit): %.1f\n", earningAfterTax)
	fmt.Printf("EBT/profit ratio: %.1f\n", ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}

func calculateFinancial(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningBeforeTax := revenue - expenses
	earningAfterTax := earningBeforeTax * (1 - taxRate/100)
	ratio := earningBeforeTax / earningAfterTax

	return earningBeforeTax, earningAfterTax, ratio
}
