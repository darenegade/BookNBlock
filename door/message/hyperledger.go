package message

import (
	"github.com/darenegade/BookNBlock/door"
)

type (
	Hyperledger struct {
		// ...?
	}
)

func NewHyperledger() *Hyperledger {
	// ??
	panic("not yet implemented")
}

func (w *Hyperledger) Subscribe() (<-chan door.OpenDoorMessage, error) {
	panic("not yet implemented")
}
