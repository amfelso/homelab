package main

import (
	"flag"
	"log"
	"time"

	collectors "github.com/amfelso/homelab/pi-agent/collectors"
)

func main() {
	var seconds int

	flag.IntVar(&seconds, "seconds", 300, "the time between metric collection cycles")

	flag.Parse()
	log.Printf("Seconds: %d\n", seconds)

	metrics := []collectors.Collector{&collectors.Cpu{}}

	ticker := time.NewTicker(time.Duration(seconds) * time.Second)
	for range ticker.C {
		for _, c := range metrics {
			val, err := c.Collect()
			if err != nil {
				log.Printf("Couldn't get metric value: %s\n", err)
				continue
			}
			log.Printf("Found metric value: %.2f\n", val)
		}
	}
}
