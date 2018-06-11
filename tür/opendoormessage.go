package tür

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"strconv"
	"strings"
)

type (
	OpenDoorMessage struct {
		DoorID       DoorID
		RenterPubkey RenterPubkey
		Timestamp    int
	}

	OpenDoorMessageHyperledger struct {
		DoorID       DoorID
		Payload      []byte
		RenterPubkey RenterPubkey
		Timestamp    int
	}
)

func (msg *OpenDoorMessageHyperledger) Decrypt(pemString string) {
	var tmp int64
	block, _ := pem.Decode([]byte(pemString))
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	decryptText, err := rsa.DecryptPKCS1v15(nil, key, msg.Payload)
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(decryptText), "s")

	if len(data) != 2 {
		panic("Reviced Paylod is invalid")
	}
	tmp, err = strconv.ParseInt(data[0], 10, 64)
	msg.Timestamp = int(tmp)
	if err != nil {
		panic(err)
	}

	msg.RenterPubkey = RenterPubkey(data[1])
}
