package message

import (
	"testing"

	"fmt"

	"github.com/darenegade/BookNBlock/door"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestWhisper_Subscribe(t *testing.T) {

	w := StartNode(WhisperConfig{})

	const hexPrivateKey = `6ecd6756d5e9d9df44be83b82d99b17983ae5ce9d0f2de9dcd68c80197aafc4a`
	const hexPublicKey = `04f0f871df7b11b3a186210ef251d10837ccfb757de9d8669225bcf73632853def72ae7680f8acdfa1ac94345017d2b4c185275a1ea2f7bbe03e939146ba355889`

	privateKey, err := crypto.HexToECDSA(hexPrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	c, err := w.Subscribe(door.DoorPrivateKey(hexPrivateKey))
	if err != nil {
		t.Fatal(err)
	}
	err = w.Post(door.OpenDoorMessage{
		DoorPublicKey: door.DoorPublicKey(hexPublicKey),
	}, privateKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s hello", <-c)
}
