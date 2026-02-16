package main

import (
	"fmt"

	// "example.com/price-calculator/cmd_manager"
	"example.com/price-calculator/file_manager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.10, 0.15}
	doneChannels := make([]chan bool, len(taxRates))
	errChannels := make([]chan error, len(taxRates))

	for idx, taxRate := range taxRates {
		doneChannels[idx] = make(chan bool)
		errChannels[idx] = make(chan error)

		fileManager := file_manager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdManager := cmd_manager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fileManager, taxRate)
		go priceJob.Process(doneChannels[idx], errChannels[idx])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	for idx := range taxRates {
		select {
		case err := <-errChannels[idx]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChannels[idx]:
			fmt.Println("Done!!")
		}
	}

}
