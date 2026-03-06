package main

import (
	"flag"
	"log"
	"time"

	collectors "github.com/amfelso/homelab/pi-agent/collectors"
)

func main() {
	var schedule int

	flag.IntVar(&schedule, "schedule", 300, "seconds between metric collection cycles")

	flag.Parse()
	log.Printf("Schedule: %d seconds\n", schedule)

	metrics := []collectors.Collector{collectors.NewCpu("/host/proc/stat"), collectors.NewMemory("/host/proc/meminfo"),
									  collectors.NewTemperature("/host/sys/class/thermal/thermal_zone0/temp")}

	ticker := time.NewTicker(time.Duration(schedule) * time.Second)
	for range ticker.C {
		for _, c := range metrics {
			val, err := c.Collect()
			if err != nil {
				log.Printf("Couldn't get metric value: %s\n", err)
				continue
			}
			log.Printf("%s: %.2f\n", c.Name(), val)
		}
	}
}
