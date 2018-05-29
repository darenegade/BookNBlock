package mqtt

import (
	"sync"
	"testing"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Connection(t *testing.T) {

	const TOPIC = "test"

	opts := mqtt.NewClientOptions().AddBroker("tcp://104.196.103.14:1883")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}

	var wg sync.WaitGroup
	wg.Add(1)

	if token := client.Subscribe(TOPIC, 0, func(client mqtt.Client, msg mqtt.Message) {
		if string(msg.Payload()) != "mymessage" {
			t.Fatalf("want mymessage, got %s", msg.Payload())
		}

		// Hier muss unsere Aufsperrcode rein
		// 1. Auslesen der Nachricht

		// 2. Validierung der Nachricht
		// 3. Öffnen dert Tür

		wg.Done()
	}); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}

	if token := client.Publish(TOPIC, 0, false, "mymessage"); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}
	wg.Wait()
}
