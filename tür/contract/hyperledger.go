package contract

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	time "time"
)

type (
	HyperLedger struct {
		Url   string
		offer Offer
	}
	Offer struct {
		OfferID    int64     `json:"offerId"`
		Free       bool      `json:"free"`
		Price      float64   `json:"price"`
		CheckIn    time.Time `json:"checkIn"`
		CheckOut   time.Time `json:"checkOut"`
		ObjectName string    `json:"objectName"`
		OwnerName  string    `json:"ownerName"`
		TenantPk   string    `json:"tenantPk"`
		LandlordPk string    `json:"landlordPk"`
	}
)

//renterID is not needed here (in contrast to the equivalent method for ethereum), since decrypting the message with the PK is proof enough for the requester's authenticity
func (h *HyperLedger) isAllowedAt(requestPointofTime time.Time, startTime time.Time, endTime time.Time) (allowed bool) {
	if requestPointofTime.Before(endTime) && requestPointofTime.After(startTime) {
		return true
	}
	return false
}

func (h *HyperLedger) getBlockData() (offer Offer) {
	fmt.Println("Starting the application...")
	response, err := http.Get(h.Url)
	var responseData []byte
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		responseData, err = ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(responseData))
	}
	offer = Offer{} // Slice of Offer instances
	json.Unmarshal(responseData, &offer)
	return offer
}
func (h *HyperLedger) getRenterID(payload string) {
	// if h.offer == nil {
	// 	return
	// }
}

//Needed for decrypting the encrypted message from the user (over the mqqt broker)
func (h *HyperLedger) decryptPayload(renterPK string, encryptedPayload string) (decryptedPayload string) {
	return
}

// 	key, _ := hex.DecodeString(renterPK)
// 	ciphertext, _ := hex.DecodeString(encryptedPayload)
// 	return
