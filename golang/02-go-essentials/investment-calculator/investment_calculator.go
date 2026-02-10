package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmentAmount float64
	var years float64

	expectedReturnRate := 5.5

	// fmt.Print("Enter investment amount: ")
	outputText("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Enter expected return rate: ")
	outputText("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Enter investment duration: ")
	outputText("Enter investment duration: ")
	fmt.Scan(&years)

	// futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// formattedFutureValue := fmt.Sprintf("Future Value: %.1f\n", futureRealValue)
	// formattedFutureRealValue := fmt.Sprintf("Future Value (adjusted for inflation): %.1f\n", futureRealValue)
	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Value (adjusted for inflation):", futureRealValue)

	fmt.Printf("Future Value: %.0f\nFuture Value (adjusted for inflation): %.0f\n", futureValue, futureRealValue)

	// fmt.Printf(`Future Value: %.0f
	// 	Future Value (adjusted for inflation): %.0f`,futureValue, futureRealValue)

	// fmt.Println("Future Value (adjusted for inflation)", futureRealValue)
	// fmt.Print(formattedFutureValue, formattedFutureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (futureValue float64, realFutureValue float64) {
	futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	realFutureValue = futureValue / math.Pow(1+inflationRate/100, years)

	return futureValue, realFutureValue
	// return
}
