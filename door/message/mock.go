package message

import (
	"time"

	"github.com/darenegade/BookNBlock/door"
)

type (
	Mock struct{}
)

func (Mock) Subscribe() (<-chan door.OpenDoorMessage, error) {
	c := make(chan door.OpenDoorMessage)
	go func() {
		for {
			c <- door.OpenDoorMessage{
				RenterPubkey: door.RenterPubkey("franz"),
			}
			time.Sleep(time.Second * 10)
		}
	}()
	return c, nil
}
