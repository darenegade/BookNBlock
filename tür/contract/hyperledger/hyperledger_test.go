package hyperledger

import (
	"fmt"
	"io/ioutil"
	http "net/http"
	"os"
	"testing"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestMain(m *testing.M) {
	//Setting up the HTTP mock to have a fake rest interface
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	fixture := map[string]interface{}{"offerId": 123, "free": false, "price": 1200, "checkIn": "2018-01-01T00:00:00Z", "checkOut": "2019-01-01T00:00:00Z", "objectName": "Test-Apartment", "ownerName": "Schlicht", "tenantPk": "MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAId84WY3SCyOwh602wC0OTxQmuPZ3MwfIeNbnhYBs2Pnx/eq5KO0Mh5Bu6X0sGBGHS47Kd1bZ7GJgAFvGWLHr+kCAwEAAQ", "landlordPk": "test"}

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
	fixture := `{"checkIn":"2018-01-01T00:00:00Z","checkOut":"2019-01-01T00:00:00Z","free":false,"landlordPk":"test","objectName":"Test-Apartment","offerId":123,"ownerName":"Schlicht","price":1200,"tenantPk":"MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAId84WY3SCyOwh602wC0OTxQmuPZ3MwfIeNbnhYBs2Pnx/eq5KO0Mh5Bu6X0sGBGHS47Kd1bZ7GJgAFvGWLHr+kCAwEAAQ"}`
	resp, err := http.Get(fakeURL)
	if err != nil {
		t.Error("Test api not working.")

	}
	responseData, err := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(responseData))
	fmt.Print("\n\n")
	fmt.Printf((string(fixture)))
	if string(responseData) == fixture {
		t.Error("Test api returns wrong json.")
	}

}

func TestGetBlockData(t *testing.T) {
	fakeURL := "https://hyperledger.com/door/123"
	hyper := HyperLedger{URL: fakeURL}
	fmt.Print("asdfasdf" + hyper.URL)
	hyper.getBlockData()
	fmt.Print("it workded!!!")
	fmt.Print(hyper.offer)
	// if hyper.offer.CheckIn != time.Parse("2018-01-01T00:00:00Z", "2018-01-01T00:00:00Z") {
	// 	t.Error("getBlockData not working.")
	// }
}
