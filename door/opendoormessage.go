package door

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"strconv"
	"strings"
)

type (
	OpenDoorMessage struct {
		DoorID       DoorPublicKey
		RenterPubkey RenterPublicKey
		Timestamp    int
		Booking      BookingID
	}

	OpenDoorMessageHyperledger struct {
		DoorID       DoorPublicKey
		Payload      []byte
		RenterPubkey RenterPublicKey
		Timestamp    int
		Booking      BookingID
	}
)

func (msg *OpenDoorMessageHyperledger) Decrypt(pemString string) {
	var tmp int64
	block, _ := pem.Decode([]byte(pemString))

	key, err2 := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err2 != nil {
		panic(err2)
	}

	decryptText, err := rsa.DecryptPKCS1v15(rand.Reader, key, msg.Payload)
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(decryptText), ",")

	if len(data) != 2 {
		panic("Reviced Paylod is invalid")
	}
	tmp, err = strconv.ParseInt(data[0], 10, 64)
	msg.Timestamp = int(tmp)
	if err != nil {
		panic(err)
	}

	msg.RenterPubkey = RenterPublicKey(data[1])
}
