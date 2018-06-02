package message

import (
	"fmt"

	".."
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const TOPIC = "test"

type (
	Hyperledger struct {
		client mqtt.Client
	}
)

func NewHyperledger() *Hyperledger {

	opts := mqtt.NewClientOptions().AddBroker("tcp://104.196.103.14:1883")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error)
	}

	return &Hyperledger{
		client: client,
	}

}

func (h *Hyperledger) Subscribe() (<-chan tür.OpenDoorMessage, error) {
	c := make(chan tür.OpenDoorMessage)
	if token := h.client.Subscribe(TOPIC, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Println(msg.Payload())

		c <- tür.OpenDoorMessage{
			//MieterID: "asdf",
		}

	}); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return c, nil
}
