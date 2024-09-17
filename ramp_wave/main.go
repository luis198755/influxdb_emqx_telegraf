package main

import (
	"fmt"
	"math"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	topic     = "sensor/ramp"
	broker    = "tcp://emqx:1883"
	clientID  = "go-ramp-simulator"
	amplitude = 1.0
	period    = 10.0 // seconds
)

func rampWave(t float64) float64 {
	return amplitude * (2 * (t/period - math.Floor(0.5+t/period)))
}

func main() {
	opts := mqtt.NewClientOptions().AddBroker(broker)
	opts.SetClientID(clientID)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Connection error:", token.Error())
		os.Exit(1)
	}

	t := 0.0
	for {
		value := rampWave(t)
		message := fmt.Sprintf("%.2f", value)

		token := client.Publish(topic, 0, false, message)
		token.Wait()

		//fmt.Printf("Published: %s to topic: %s\n", message, topic)

		time.Sleep(100 * time.Millisecond)
		t += 0.1
		if t >= period {
			t = 0
		}
	}
}
