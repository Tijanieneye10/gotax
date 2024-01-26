package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {

	var floats []float64
	for _, line := range strings {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, errors.New("failed to convert text to float")
		}

		floats = append(floats, floatPrice)
	}

	return floats, nil
}
