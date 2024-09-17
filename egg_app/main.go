package main

import (
	"fmt"
	"math"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	topic    = "sensor/eeg"
	broker   = "tcp://emqx:1883"
	clientID = "go-eeg-simulator"
)

// Simula una señal EEG combinando diferentes ondas cerebrales
func simulateEEG(t float64) float64 {
	// Simula ondas delta (1-4 Hz)
	delta := math.Sin(2*math.Pi*2*t) * 30

	// Simula ondas theta (4-8 Hz)
	theta := math.Sin(2*math.Pi*6*t) * 20

	// Simula ondas alpha (8-13 Hz)
	alpha := math.Sin(2*math.Pi*10*t) * 15

	// Simula ondas beta (13-30 Hz)
	beta := math.Sin(2*math.Pi*20*t) * 10

	// Combina las ondas y añade un poco de ruido aleatorio
	eeg := delta + theta + alpha + beta + (math.Sin(t*100) * 5)

	return eeg
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
		value := simulateEEG(t)
		message := fmt.Sprintf("%.2f", value)

		token := client.Publish(topic, 0, false, message)
		token.Wait()
		fmt.Printf("Publicado: %s en el tema: %s\n", message, topic)

		t += 0.01 // Incrementa el tiempo (asumiendo 100 muestras por segundo)
		time.Sleep(1000 * time.Millisecond)
	}
}
