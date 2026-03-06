package main

import (
	"flag"
	"log"
	"os"
	"time"

	collectors "github.com/amfelso/homelab/pi-agent/collectors"
	publisher "github.com/amfelso/homelab/pi-agent/publisher"
)

func main() {
	var schedule int

	// Set schedule
	flag.IntVar(&schedule, "schedule", 300, "seconds between metric collection cycles")
	flag.Parse()
	log.Printf("Schedule: %d seconds\n", schedule)

	// Load env
	nodeName := os.Getenv("NODE_NAME")
	mqttUser := os.Getenv("MQTT_USERNAME")
	mqttPwd := os.Getenv("MQTT_PASSWORD")
	mqttBroker := os.Getenv("MQTT_BROKER_URL")
	log.Printf("Node: %s\n", nodeName)

	// Create collectors and publisher
	metrics := []collectors.Collector{collectors.NewCpu("/host/proc/stat"), collectors.NewMemory("/host/proc/meminfo"),
		collectors.NewTemperature("/host/sys/class/thermal/thermal_zone0/temp"),
		collectors.NewDisk("/host/root/")}
	p, err := publisher.NewPublisher(mqttBroker, nodeName, "/ca.crt", mqttUser, mqttPwd)
	if err != nil {
		log.Fatalf("Couldn't connect to publisher: %s\n", err)
	}

	ticker := time.NewTicker(time.Duration(schedule) * time.Second)
	for range ticker.C {
		for _, c := range metrics {
			val, err := c.Collect()
			if err != nil {
				log.Printf("Couldn't get metric value: %s\n", err)
				continue
			}
			log.Printf("%s: %.2f\n", c.Name(), val)
			err = p.Publish(c.Name(), val)
			if err != nil {
				log.Printf("Couldn't publish metric value: %s\n", err)
				continue
			}
		}
	}
}
