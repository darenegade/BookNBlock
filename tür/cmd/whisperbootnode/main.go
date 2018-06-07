package main

import (
	"../.."
	"../../message"
)

// main used to start an bootnode for the whisper protocol. It is used to provide an initial
// central network point for door or renter to join.
func main() {
	c := message.WhisperConfig{
		ListenAddr: "127.0.0.1:30349",
		NodeID:     `b3651aff593ef395ee7c16f3ca681830f7d8d0b2729cf472b14f2c4ebe833aa0`,
		HTTPPort:   9945,
	}
	w := message.StartNode(c)
	w.Subscribe(tür.TürID("6ecd6756d5e9d9df44be83b82d99b17983ae5ce9d0f2de9dcd68c80197aafc4a"))
	w.Node.Wait()
}
