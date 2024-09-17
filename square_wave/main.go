package main

import (
	"fmt"
	"math"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	topic     = "sensor/square"
	broker    = "tcp://emqx:1883"
	clientID  = "go-square-simulator"
	amplitude = 1.0
	frequency = 0.1 // Hz
)

func squareWave(t float64) float64 {
	if math.Sin(2*math.Pi*frequency*t) >= 0 {
		return amplitude
	}
	return -amplitude
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
		value := squareWave(t)
		message := fmt.Sprintf("%.2f", value)

		token := client.Publish(topic, 0, false, message)
		token.Wait()

		//fmt.Printf("Published: %s to topic: %s\n", message, topic)

		time.Sleep(100 * time.Millisecond)
		t += 0.1
	}
}
