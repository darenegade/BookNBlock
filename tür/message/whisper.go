package message

import (
	".."
)

type (
	Whisper struct {
		// ...?
	}
)

func NewWhipser() *Whisper {
	// ??
	panic("not yet implemented")
}

func (w *Whisper) Subscribe() (<-chan tür.OpenDoorMessage, error) {
	c := make(chan tür.OpenDoorMessage)
	panic("not yet implemented")
	return c, nil
}
