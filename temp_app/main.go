package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	topic          = "sensor/temperature"
	broker         = "tcp://emqx:1883"
	clientID       = "go-temp-simulator"
	minTemperature = 15.0
	maxTemperature = 30.0
)

func main() {
	opts := mqtt.NewClientOptions().AddBroker(broker)
	opts.SetClientID(clientID)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Connection error:", token.Error())
		os.Exit(1)
	}

	for {
		temperature := minTemperature + rand.Float64()*(maxTemperature-minTemperature)
		message := fmt.Sprintf("%.2f", temperature)

		token := client.Publish(topic, 0, false, message)
		token.Wait()

		//fmt.Printf("Published: %s to topic: %s\n", message, topic)

		time.Sleep(10 * time.Second)
	}
}
