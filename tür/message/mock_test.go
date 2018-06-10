package message

import (
	"testing"
	"time"

	".."
)

func TestMock_Subscribe(t *testing.T) {
	m := Mock{}
	c, _ := m.Subscribe()
	time.Sleep(time.Millisecond)
	message := <-c
	if message.RenterID != tür.RenterID("franz") {
		t.Error("Nachricht war nicht für franz :(")
	}
}
