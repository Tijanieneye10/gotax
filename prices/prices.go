package prices

import (
	"example.com/gotax/conversion"
	"example.com/gotax/iomanager"
	"fmt"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func NewTaxIncludedPriceJob(io iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		IOManager:   io,
	}
}

func (job *TaxIncludedPriceJob) Process() {

	job.loadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrices := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrices)
	}

	job.TaxIncludedPrices = result

	err := job.IOManager.WriteResult(job)

	if err != nil {
		return
	}

	fmt.Println("file saved successfully")
}

func (job *TaxIncludedPriceJob) loadData() {

	//Convert the text to float64
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return
	}
	prices, _ := conversion.StringToFloat(lines)
	job.InputPrices = prices
}
