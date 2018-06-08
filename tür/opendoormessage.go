package tür

import (
	"crypto"
	"crypto/rsa"
)

type (
	OpenDoorMessage struct {
		DoorID       DoorID
		RenterID     RenterID
		RenterPubkey RenterPubkey
		Timestamp    int64
	}

	OpenDoorMessageHyperledger struct {
		DoorID    DoorID
		Payload   []byte
		RenterID  RenterID
		Timestamp int64
	}
)

func (msg *OpenDoorMessageHyperledger) Encrypt(pub *rsa.PublicKey, sig []byte) {
	rsa.VerifyPKCS1v15(pub, crypto.SHA256, msg.Payload, sig)
}
