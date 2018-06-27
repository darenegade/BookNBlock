package message

import (
	"strings"
	"testing"
	"time"
)

func TestHyperledger_Subscribe(t *testing.T) {
	config := Config{BrokerAdress: "localhost:1883"}
	h := NewHyperledger(config)

	h.SendtestMessage()

	c, _ := h.Subscribe()
	h.SendtestMessage()

	time.Sleep(200 * time.Millisecond)

	message := <-c
	if message.DoorID != "008457" {
		t.Error("DoorId war nicht 008457.")
	}
	if string(message.Payload) != "5fb72a1c804906d8d4d799e6d2fa7414085fc4d3687e462ac256f18c0dfe06f5d79a70" {
		t.Error("Payload ist nicht 5fb72a1c804906d8d4d799e6d2fa7414085fc4d3687e462ac256f18c0dfe06f5d79a70.")
	}

}

func TestEncypt(t *testing.T) {
	config := Config{BrokerAdress: "localhost:1883"}
	h := NewHyperledger(config)
	msg := h.SendtestMessage()
	if strings.Compare(msg, "{ \"doorID\": \"008457\", \"payload\": \"5fb72a1c804906d8d4d799e6d2fa7414085fc4d3687e462ac256f18c0dfe06f5d79a70\" }") == 0 {
		t.Error("Test message funktioniert nicht")
	}
}
