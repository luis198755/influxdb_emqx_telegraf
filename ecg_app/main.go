package main

import (
	"fmt"
	"math"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	topic    = "sensor/ecg"
	broker   = "tcp://emqx:1883"
	clientID = "go-ecg-simulator"
)

func ecgWave(t float64) float64 {
	p := -5.0 * math.Pow(t, 3) * math.Exp(-t)
	q := 15.0 * math.Pow(t, 2) * math.Exp(-2*t)
	r := 0.3 * math.Exp(-0.5*math.Pow(t-2, 2))
	s := 3.0 * math.Pow(t-2, 4) * math.Exp(-3*(t-2))
	return p + q + r + s
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
		value := ecgWave(math.Mod(t, 5.0))
		message := fmt.Sprintf("%.2f", value)

		token := client.Publish(topic, 0, false, message)
		token.Wait()

		//fmt.Printf("Published: %s to topic: %s\n", message, topic)

		time.Sleep(100 * time.Millisecond)
		t += 0.1
	}
}
