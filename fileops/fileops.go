package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadFLoatFromFile(fileName string, defaultValue float64) (float64, error) {
	content, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValue, errors.New("file not found")
	}

	valuetext := string(content)
	value, error := strconv.ParseFloat(valuetext, 64)

	if error != nil {
		return defaultValue, errors.New("Failed to parse")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valuetext := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valuetext), 0644)
}
