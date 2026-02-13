package main

import (
	"fmt"

	"example.com/price-calculator/cmd_manager"
	// "example.com/price-calculator/file_manager"
	"example.com/price-calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.07, 0.10, 0.15}

	for _, taxRate := range taxRates {
		// fileManager := file_manager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdManager := cmd_manager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdManager, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}

}
