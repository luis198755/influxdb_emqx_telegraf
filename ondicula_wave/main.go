package main

import (
	"fmt"
	"math"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	topic    = "sensor/wavelet"
	broker   = "tcp://emqx:1883"
	clientID = "go-wavelet-generator"
)

func wavelet(t float64) float64 {
	// Wavelet de Morlet simplificada
	return math.Exp(-t*t/2) * math.Cos(5*t)
}

func main() {
	opts := mqtt.NewClientOptions().AddBroker(broker)
	opts.SetClientID(clientID)
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Error de conexión:", token.Error())
		os.Exit(1)
	}

	t := 0.0
	for {
		value := wavelet(t)
		message := fmt.Sprintf("%.6f", value)

		token := client.Publish(topic, 0, false, message)
		token.Wait()
		fmt.Printf("Publicado: %s en el tema: %s\n", message, topic)

		t += 0.1
		if t > 10 {
			t = 0
		}

		time.Sleep(100 * time.Millisecond)
	}
}