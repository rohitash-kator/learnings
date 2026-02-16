package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/io_manager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64              `json:"tax_rate"`
	InputPrices       []float64            `json:"input_prices"`
	TaxIncludedPrices map[string]string    `json:"tax_included_prices"`
	IOManager         io_manager.IOManager `json:"-"` // - is used to ignore the field from the JSON output
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.Read()

	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices

	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChannel chan bool, errChannel chan error) {
	err := job.LoadData()

	if err != nil {
		errChannel <- err
		return
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	job.IOManager.Write(job)

	doneChannel <- true
}

func NewTaxIncludedPriceJob(ioManager io_manager.IOManager, taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		IOManager:   ioManager,
	}
}
