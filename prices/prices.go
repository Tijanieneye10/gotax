package prices

import (
	"bufio"
	"example.com/gotax/conversion"
	"fmt"
	"os"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process() {

	job.loadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrices := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrices)
	}

	fmt.Println(result)
}

func (job *TaxIncludedPriceJob) loadData() {

	file, err := os.Open("price.txt")

	if err != nil {
		fmt.Println("An error occurred")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	//Convert the text to float64
	prices, _ := conversion.StringToFloat(lines)
	job.InputPrices = prices
	_ = file.Close()
}
