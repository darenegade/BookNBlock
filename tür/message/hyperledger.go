package message

import (
	"encoding/json"
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
	// localhost
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")

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
		var dat map[string]interface{}
		fmt.Println(msg.Payload())

		if err := json.Unmarshal(msg.Payload(), &dat); err != nil {
			panic(err)
		}
		fmt.Println(dat)

		c <- tür.OpenDoorMessage{
			RenterID:  tür.RenterID(dat["renterID"].(string)),
			RenterPK:  tür.RenterPK(dat["renterPK"].(string)),
			DoorID:    tür.DoorID(dat["doorID"].(string)),
			Timestamp: int(dat["timestamp"].(float64)),
		}

	}); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return c, nil
}

func (h *Hyperledger) SendtestMessage() {
	testMsg := "{ \"renterID\": \"4286f4\", \"doorID\": \"008457\", \"renterPK\": \"f78uf\", \"timestamp\":1527950669609 }"
	if token := h.client.Publish(TOPIC, 0, false, testMsg); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
