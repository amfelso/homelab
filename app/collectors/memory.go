package collectors

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Memory struct{
	fileName string
}

func NewMemory(path string) *Memory {
	return &Memory{fileName: path}
}
 
func (m *Memory)Name() string {
    return "memory"
}

func (m *Memory)Collect() (float64, error) {
	// Open the file
	file, err := os.Open(m.fileName)
	if err != nil {
		log.Printf("Error opening file: %s", err)
		return 0.0, err 
	}

	// Ensure the file is closed the the function exits
	defer func() { _ = file.Close() }()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	
	// Get memTotal from line 0
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return 0.0, err
		}
		return 0.0, fmt.Errorf("unexpected EOF in %s", m.fileName)
	}
	memTotal, err := strconv.ParseUint(strings.Fields(scanner.Text())[1], 10, 64)
	if err != nil {
		return 0.0, err
	}

	// Discard line 1 (MemFree)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return 0.0, err
		}
		return 0.0, fmt.Errorf("unexpected EOF in %s", m.fileName)
	}

	// Get memAvailable from line 2
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return 0.0, err
		}
		return 0.0, fmt.Errorf("unexpected EOF in %s", m.fileName)
	}
	memAvailable, err := strconv.ParseUint(strings.Fields(scanner.Text())[1], 10, 64)
	if err != nil {
		return 0.0, err
	}

	// Calculate usagePercent
	usagePercent := float64(memTotal - memAvailable) / float64(memTotal) * 100
	return usagePercent, nil
}