package main

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func OnMessageReceived(client MQTT.Client, message MQTT.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
}

func main() {
	println("Main Function")
}
