package tür

import (
	"crypto/rsa"
	"strconv"
	"strings"
)

type (
	OpenDoorMessage struct {
		DoorID       DoorID
		RenterPubkey RenterPubkey
		Timestamp    int64
	}

	OpenDoorMessageHyperledger struct {
		DoorID       DoorID
		Payload      []byte
		RenterPubkey RenterPubkey
		Timestamp    int64
	}
)

func (msg *OpenDoorMessageHyperledger) Decrypt(pub *rsa.PrivateKey) {
	decryptText, err := rsa.DecryptPKCS1v15(nil, pub, msg.Payload)
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(decryptText), "s")

	if len(data) != 2 {
		panic("Reviced Paylod is invalid")
	}
	msg.Timestamp, err = strconv.ParseInt(data[0], 10, 64)
	if err != nil {
		panic(err)
	}

	msg.RenterPubkey = RenterPubkey(data[1])
}
