package message

import (
	"testing"
	"time"
)

func TestHyperledger_Subscribe(t *testing.T) {
	h := NewHyperledger()

	h.SendtestMessage()

	c, _ := h.Subscribe()
	h.SendtestMessage()

	time.Sleep(200 * time.Millisecond)

	message := <-c
	if message.RenterID != "4286f4" {
		t.Error("RenterID war nicht 4286f4.")
	}
	if message.RenterPK != "f78uf" {
		t.Error("RenterPK war nicht f78uf.")
	}
	if message.DoorID != "008457" {
		t.Error("DoorId war nicht 008457.")
	}
	if message.Timestamp != 1527950669609 {
		t.Error("Timestamp ist nicht 1527950669609.")
	}

}

func TestEncypt(t *testing.T) {
	h := NewHyperledger()
	msg := h.SendtestMessage()
	if msg != "{ \"doorID\": \"008457\", \"renterID\": \"4286f4\", \"payload\": \"5fb72a1c804906d8d4d799e6d2fa7414085fc4d3687e462ac256f18c0dfe06f5d79a70\" }" {
		t.Error("Test message funktioniert nicht")
	}
}
