package hyperledger

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
		URL   string
		offer *Offer
	}
	//In this case OfferID corresponds to DoorID
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

func (h *HyperLedger) getBlockData() {
	fmt.Println("Starting the application...")
	response, err := http.Get(h.URL)
	var responseData []byte
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return
	} else {
		responseData, err = ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
			return
		}

	}
	var offer Offer
	err = json.Unmarshal(responseData, &offer)
	if err != nil {
		fmt.Print("Unmarshalling did not work")
		log.Fatal(err)
		return
	}

	h.offer = &offer

}
