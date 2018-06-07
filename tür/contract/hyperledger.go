package contract

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	time "time"
)

type (
	HyperLedger struct {
		Url string
	}
)

//renterID is the users Public Key
func (h *HyperLedger) isAllowedAt(renterID string, requestPointofTime time.Time, startTime time.Time, endTime time.Time) (allowed bool) {

	return false
}

func (h *HyperLedger) getBlockData() (responseData []byte) {
	fmt.Println("Starting the application...")
	response, err := http.Get(h.Url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		responseData, err = ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(responseData))
	}

	return responseData
}

package contract

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	time "time"
)

type (
	HyperLedger struct {
		Url string
	}
)

//renterID is the users Public Key
func (h *HyperLedger) isAllowedAt(renterID string, requestPointofTime time.Time, startTime time.Time, endTime time.Time) (allowed bool) {

	return false
}

func (h *HyperLedger) getBlockData() (responseData []byte) {
	fmt.Println("Starting the application...")
	response, err := http.Get(h.Url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
	return responseData
}

func (h *Hyperledger) decryptPaylpad(payload string) (renterPK tür.RenterPK, timestamp int64) {
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
 
