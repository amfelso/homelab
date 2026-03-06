package collectors

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Cpu struct {
	prevIdle  uint64
	prevTotal uint64
	fileName  string
}

func NewCpu(path string) *Cpu {
	return &Cpu{fileName: path}
}

func (c *Cpu) Name() string {
	return "cpu"
}

func (c *Cpu) Collect() (float64, error) {
	// Open the file
	file, err := os.Open(c.fileName)
	if err != nil {
		log.Printf("Error opening file: %s", err)
		return 0.0, err
	}

	// Ensure the file is closed the the function exits
	defer func() { _ = file.Close() }()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return 0.0, err
		}
		return 0.0, fmt.Errorf("unexpected EOF in %s", c.fileName)
	}

	stats := strings.Fields(scanner.Text())[1:]

	// Caculate currentIdle
	currentIdle, err := strconv.ParseUint(stats[3], 10, 64) // idle is index 3
	if err != nil {
		return 0.0, err
	}

	// Calculate idleDelta
	idleDelta := currentIdle - c.prevIdle

	// calculate currentTotal
	var currentTotal uint64
	for _, stat := range stats {
		stat, err := strconv.ParseUint(stat, 10, 64)
		if err != nil {
			return 0.0, err
		}
		currentTotal += uint64(stat)
	}

	// values for first run
	if c.prevTotal == 0 {
		c.prevIdle = currentIdle
		c.prevTotal = currentTotal
		return 0.0, nil
	}

	// calculate totalDelta
	totalDelta := currentTotal - c.prevTotal

	// Usage Percent = (1 - idleDelta/totalDelta) * 100
	usagePercent := (1.0 - float64(idleDelta)/float64(totalDelta)) * 100

	// Save current readings
	c.prevIdle = currentIdle
	c.prevTotal = currentTotal

	return usagePercent, nil
}
