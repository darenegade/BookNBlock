package contract

import (
	"fmt"
	"io/ioutil"
	http "net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
)

var fixture string

func setUp() {
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
	fmt.Print("dsfasdfasdfasdfasdf")
}
func TestMain(m *testing.M) {
	setUp()

	retCode := m.Run()
	os.Exit(retCode)

}
func TestMockAPI(t *testing.T) {
	fakeURL := "https://hyperledger.com/door/123"
	resp, err := http.Get(fakeURL)
	if err != nil {
		t.Error("Test api not working.")
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	fmt.Print(string(responseData))
	fmt.Print(fixture)
	if string(responseData) != fixture {
		t.Error("Test api returns wrong json.")
	}

}
