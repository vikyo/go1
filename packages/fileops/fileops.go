package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatValueToFile(value float64, fileName string) {
	os.WriteFile(fileName, []byte(fmt.Sprint(value)), 0644)
}

func ReadFloatValueFromFile(filePath string) (float64, error) {
	arg, err := os.ReadFile(filePath)

	if err != nil {
		return 1000, errors.New("failed to read file")
	}
	valueText := string(arg)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse value from file")
	}

	return value, nil
}
