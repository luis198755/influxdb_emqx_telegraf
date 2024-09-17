package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	topic    = "sensor/gaussian_noise"
	broker   = "tcp://emqx:1883"
	clientID = "go-gaussian-noise-simulator"
	mean     = 0.0
	stdDev   = 1.0
)

func gaussianNoise() float64 {
	return rand.NormFloat64()*stdDev + mean
}

func main() {
	rand.Seed(time.Now().UnixNano())

	opts := mqtt.NewClientOptions().AddBroker(broker)
	opts.SetClientID(clientID)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Connection error:", token.Error())
		os.Exit(1)
	}

	for {
		value := gaussianNoise()
		message := fmt.Sprintf("%.6f", value)

		token := client.Publish(topic, 0, false, message)
		token.Wait()

		//fmt.Printf("Published: %s to topic: %s\n", message, topic)

		time.Sleep(100 * time.Millisecond)
	}
}
