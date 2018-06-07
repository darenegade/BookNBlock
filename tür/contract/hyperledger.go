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
