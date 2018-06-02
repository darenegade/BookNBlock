package message

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	".."
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const TOPIC = "test"

type (
	Hyperledger struct {
		client mqtt.Client
	}
)

func NewHyperledger() *Hyperledger {
	// localhost
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error)
	}

	return &Hyperledger{
		client: client,
	}

}

func (h *Hyperledger) Subscribe() (<-chan tür.OpenDoorMessage, error) {
	c := make(chan tür.OpenDoorMessage)
	if token := h.client.Subscribe(TOPIC, 0, func(client mqtt.Client, msg mqtt.Message) {
		var dat map[string]interface{}
		fmt.Println(msg.Payload())

		if err := json.Unmarshal(msg.Payload(), &dat); err != nil {
			panic(err)
		}
		fmt.Println(dat)

		renterID, timestamp := h.decryptPaylpad(dat["payload"].(string))

		c <- tür.OpenDoorMessage{
			DoorID:    tür.DoorID(dat["doorID"].(string)),
			RenterPK:  tür.RenterPK(dat["renterPK"].(string)),
			RenterID:  renterID,
			Timestamp: timestamp,
			//Timestamp: int(dat["timestamp"].(float64)),
		}

	}); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return c, nil
}

func (h *Hyperledger) decryptPaylpad(payload string) (renterID tür.RenterID, timestamp int64) {
	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
	ciphertext, _ := hex.DecodeString(payload)
	nonce := []byte("64a9433eae7c")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%s\n", plaintext)
	words := strings.Split(string(plaintext), ",")
	renterID = tür.RenterID(words[0])

	i, err := strconv.ParseInt(words[1], 10, 64)
	if err != nil {
		panic(err)
	}
	timestamp = i
	return
}

func (h *Hyperledger) SendtestMessage() (testMsg string) {
	testMsg = fmt.Sprintf("{ \"doorID\": \"008457\", \"renterPK\": \"f78uf\", \"payload\": \"%x\" }", h.test_encrypt())
	if token := h.client.Publish(TOPIC, 0, false, testMsg); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return
}

func (h *Hyperledger) test_encrypt() (ciphertext []byte) {
	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
	plaintext := []byte("4286f4,1527950669609")

	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err.Error())
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := []byte("64a9433eae7c")
	aesgcm, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	ciphertext = aesgcm.Seal(nil, nonce, plaintext, nil)

	fmt.Printf("Chipretext: %x\n", ciphertext)
	return
}
