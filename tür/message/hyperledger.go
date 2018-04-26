package message

import (
	".."
)

type (
	Hyperledger struct {
		// ...?
	}
)

func NewHyperledger() *Whisper {
	// ??
	panic("not yet implemented")
}

func (w *Hyperledger) Subscribe() (<-chan tür.OpenDoorMessage, error) {
	c := make(chan tür.OpenDoorMessage)
	panic("not yet implemented")
	return c, nil
}
