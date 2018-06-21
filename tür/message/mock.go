package message

import (
	"time"

	".."
)

type (
	Mock struct{}
)

func (Mock) Subscribe() (<-chan tür.OpenDoorMessage, error) {
	c := make(chan tür.OpenDoorMessage)
	go func() {
		for {
			c <- tür.OpenDoorMessage{
				RenterID: tür.RenterID("franz"),
			}
			time.Sleep(time.Second * 10)
		}
	}()
	return c, nil
}
