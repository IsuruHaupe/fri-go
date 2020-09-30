package paho

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func createClientOptions(brokerURI string, clientID string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURI)
	// opts.SetUsername(user)
	// opts.SetPassword(password)
	opts.SetClientID(clientID)
	return opts
}

// Connect : connect mqtt client to mqtt broker
func Connect(brokerURI string, clientID string) mqtt.Client {
	fmt.Println("Trying to connect (" + brokerURI + ", " + clientID + ")...")
	opts := createClientOptions(brokerURI, clientID)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Connected to server\n")
	}
	return client
}
