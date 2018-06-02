package message

import (
	"testing"
	"time"
)

func TestHyperledger_Subscribe(t *testing.T) {
	h := NewHyperledger()
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
