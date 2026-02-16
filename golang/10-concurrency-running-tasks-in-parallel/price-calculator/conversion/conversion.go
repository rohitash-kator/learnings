package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(stringValues []string) ([]float64, error) {

	var floats []float64

	for _, stringVal := range stringValues {
		floatVal, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, errors.New("Failed to convert sting to float.")
		}

		floats = append(floats, floatVal)
	}

	return floats, nil
}
