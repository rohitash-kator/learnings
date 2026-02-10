package main

import (
	"errors"
	"fmt"
	"os"
)

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

	revenue, revenueErr := getUserInput("Enter revenue: ")

	if revenueErr != nil {
		fmt.Println(revenueErr)
		return
		// panic(revenueErr)
	}

	expenses, expensesErr := getUserInput("Enter expenses: ")

	if expensesErr != nil {
		fmt.Println(expensesErr)
		return
	}

	taxRate, taxRateErr := getUserInput("Enter tax rate: ")

	if taxRateErr != nil {
		fmt.Println(taxRateErr)
		return
	}

	// if revenueErr != nil || expensesErr != nil || taxRateErr != nil {
	// 	fmt.Println()
	// 	return
	// }

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
	storeResults(earningBeforeTax, earningBeforeTax, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number")
	}

	return userInput, nil
}

func calculateFinancial(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningBeforeTax := revenue - expenses
	earningAfterTax := earningBeforeTax * (1 - taxRate/100)
	ratio := earningBeforeTax / earningAfterTax

	return earningBeforeTax, earningAfterTax, ratio
}
