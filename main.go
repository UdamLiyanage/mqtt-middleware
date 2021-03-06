package main

import (
	"fmt"
	"github.com/arangodb/go-driver"
	"os"
	"os/signal"
	"syscall"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/joho/godotenv/autoload"
)

var database driver.Database

func init() {
	databaseConnect()
}

func OnMessageReceived(_ MQTT.Client, message MQTT.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
	validateRequest(message.Payload())
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	opts := setClientOptions()
	topic := "#"

	opts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe(topic, 0, OnMessageReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to server:\n")
	}
	<-c
}

func setClientOptions() *MQTT.ClientOptions {
	opts := MQTT.NewClientOptions().AddBroker(os.Getenv("BROKER_URL"))
	opts.SetDefaultPublishHandler(OnMessageReceived)
	return opts
}
