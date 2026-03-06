package collectors

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Temperature struct{
	fileName string
}

func NewTemperature(path string) *Temperature {
	return &Temperature{fileName: path}
}
 
func (t *Temperature)Name() string {
    return "temp_f"
}

func (t *Temperature)Collect() (float64, error) {
	// Open the file
	file, err := os.ReadFile(t.fileName)
	if err != nil {
		log.Printf("Error reading file: %s", err)
		return 0.0, err 
	}
	
	tempC, err := strconv.ParseInt(strings.TrimSpace(string(file)), 10, 64)
	if err != nil {
		return 0.0, err
	}

	// Calculate Fahrenheit
	tempF := ((float64(tempC) / 1000.0) * (9.0 / 5.0)) + 32.0
	return tempF, nil
}