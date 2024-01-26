package main

import "example.com/gotax/prices"

func main() {
	taxRate := []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRate {
		var taxIncludedPrices = make([]float64, len(prices))
		prices.NewTaxIncludedPriceJob()
	}

	fmt.Println(result)
}