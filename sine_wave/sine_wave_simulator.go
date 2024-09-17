package main

import (
	"fmt"
	"math"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	topic    = "sensor/senoidal"
	broker   = "tcp://emqx:1883"
	clientID = "go-sine-simulator"
	amplitude = 10.0
	frequency = 0.1 // Hz
)

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
		value := amplitude * math.Sin(2*math.Pi*frequency*t)
		message := fmt.Sprintf("%.2f", value)
		
		token := client.Publish(topic, 0, false, message)
		token.Wait()

		//fmt.Printf("Published: %s to topic: %s\n", message, topic)
		
		time.Sleep(1 * time.Second)
		t += 1.0
	}
}