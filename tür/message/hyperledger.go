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

// TOPIC for OpenDoorMessage via Hyperledger and MQTT, to recive OpenDoorMessage from app
const TOPIC = "test"

type (
	// Hyperledger Making connection to MQTT Broker and for evaluating OpenDoorMessage
	Hyperledger struct {
		client mqtt.Client
	}
	// Config containing statick iportanten informations, like the adresse of the MQTT Broker
	Config struct {
		BrokerAdress string
	}
)

// NewHyperledger returns new Hyperledgercleint
func NewHyperledger(config Config) *Hyperledger {
	opts := mqtt.NewClientOptions().AddBroker(config.BrokerAdress)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error)
	}

	return &Hyperledger{
		client: client,
	}

}

// Subscribe listen to messages entend for this door
func (h *Hyperledger) Subscribe() (<-chan tür.OpenDoorMessage, error) {
	c := make(chan tür.OpenDoorMessage)
	if token := h.client.Subscribe(TOPIC, 0, func(client mqtt.Client, msg mqtt.Message) {
		var dat map[string]interface{}
		fmt.Println(msg.Payload())

		if err := json.Unmarshal(msg.Payload(), &dat); err != nil {
			panic(err)
		}
		fmt.Println(dat)

		// request publicKey from Hyperledgerblock

		renterPK, timestamp := h.decryptPaylpad(dat["payload"].(string), "")

		c <- tür.OpenDoorMessage{
			DoorID:    tür.DoorID(dat["doorID"].(string)),
			RenterID:  tür.RenterID(dat["renterID"].(string)),
			RenterPK:  renterPK,
			Timestamp: timestamp,
		}

	}); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return c, nil
}

func (h *Hyperledger) decryptPaylpad(payload string, publicKey string) (renterPK tür.RenterPK, timestamp int64) {
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
	renterPK = tür.RenterPK(words[0])

	i, err := strconv.ParseInt(words[1], 10, 64)
	if err != nil {
		panic(err)
	}
	timestamp = i
	return
}

// SendtestMessage for sending testmessage
func (h *Hyperledger) SendtestMessage() (testMsg string) {
	testMsg = fmt.Sprintf("{ \"doorID\": \"008457\", \"renterID\": \"4286f4\", \"payload\": \"%x\" }", h.test_encrypt())
	if token := h.client.Publish(TOPIC, 0, false, testMsg); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return
}

func (h *Hyperledger) test_encrypt() (ciphertext []byte) {
	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
	plaintext := []byte("f78uf,1527950669609")

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
