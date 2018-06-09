package contract

import (
	"fmt"
	"io/ioutil"
	http "net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestMain(m *testing.M) {
	//Setting up the HTTP mock to have a fake rest interface
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	fixture := `[{
		"offerID": 123,
		"free": false,
		"price": 300,
		"checkIn": "01.06.2018",
		"checkOut": "01.06.2020",
		"objectName": "small Apt.",
		"ownerName": "Schlicht",
		"tenantPK": "MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAId84WY3SCyOwh602wC0OTxQmuPZ3MwfIeNbnhYBs2Pnx/eq5KO0Mh5Bu6X0sGBGHS47Kd1bZ7GJgAFvGWLHr+kCAwEAAQ==",
		"landlordPK": "test"
	  }
	  ]`

	responder, err := httpmock.NewJsonResponder(200, fixture)
	if err != nil {
	}

	fakeURL := "https://hyperledger.com/door/123"
	httpmock.RegisterResponder("GET", fakeURL, responder)

	retCode := m.Run()
	os.Exit(retCode)

}
func TestMockAPI(t *testing.T) {
	fakeURL := "https://hyperledger.com/door/123"
	fixture := `[{\n\t\t\"offerID\": 123,\n\t\t\"free\": false,\n\t\t\"price\": 300,\n\t\t\"checkIn\": \"01.06.2018\",\n\t\t\"checkOut\": \"01.06.2020\",\n\t\t\"objectName\": \"small Apt.\",\n\t\t\"ownerName\": \"Schlicht\",\n\t\t\"tenantPK\": \"MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAId84WY3SCyOwh602wC0OTxQmuPZ3MwfIeNbnhYBs2Pnx/eq5KO0Mh5Bu6X0sGBGHS47Kd1bZ7GJgAFvGWLHr+kCAwEAAQ==\",\n\t\t\"landlordPK\": \"test\"\n\t  }\n\t  ]`
	resp, err := http.Get(fakeURL)
	if err != nil {
		t.Error("Test api not working.")
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(responseData))
	fmt.Print("\n\n")
	fmt.Printf((string(fixture)))
	if len(string(responseData)) == len(fixture) {
		t.Error("Test api returns wrong json.")
	}

}
