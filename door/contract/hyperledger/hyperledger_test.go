package hyperledger

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
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
	//fixture := []byte(`{"checkIn":"2018-01-01T00:00:00Z","checkOut":"2019-01-01T00:00:00Z","free":false,"landlordPk":"test","objectName":"Test-Apartment","offerId":123,"ownerName":"Schlicht","price":1200,"tenantPk":"MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAId84WY3SCyOwh602wC0OTxQmuPZ3MwfIeNbnhYBs2Pnx/eq5KO0Mh5Bu6X0sGBGHS47Kd1bZ7GJgAFvGWLHr+kCAwEAAQ"}`)
	fixture := map[string]interface{}{"offerId": 123, "free": false, "price": 1200, "checkIn": "2018-01-01T00:00:00Z", "checkOut": "2019-01-01T00:00:00Z", "objectName": "Test-Apartment", "ownerName": "Schlicht", "tenantPk": "MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAId84WY3SCyOwh602wC0OTxQmuPZ3MwfIeNbnhYBs2Pnx/eq5KO0Mh5Bu6X0sGBGHS47Kd1bZ7GJgAFvGWLHr+kCAwEAAQ", "landlordPk": "test"}

	resp, err := http.Get(fakeURL)
	if err != nil {
		t.Error("Test api not working.")

	}
	responseData, err := ioutil.ReadAll(resp.Body)

	if reflect.DeepEqual(responseData, fixture) {
		t.Error("Test api returns wrong json.")
	}

}

func TestGetBlockData(t *testing.T) {
	// fakeURL := "https://hyperledger.com/door/123"
	// hyper := HyperLedger{URL: fakeURL}
	// hyper.getBlockData()
	// if hyper.offer.OfferID != 123 {
	// 	t.Error("getBlockData not working.")
	// }
}

func TestAllowedInWithActualMainState(t *testing.T) {
	httpmock.Deactivate()
	realURL := "http://129.187.110.174:3000/api/fabric/1_0/channels/bookchannel/ledger?chaincodeId=bookandblockcc"
	hyper := HyperLedger{URL: realURL, DoorID: "1"}
	responses := hyper.getHistoryForOffer()
	key := hyper.getPubKeyOfValidUser(responses)
	allowed := hyper.isAllowedIn(key, hyper.Message)
	if !allowed {
		t.Error("isAllowedAt not working.")
	}

}

func TestConnectionToMainState(t *testing.T) {
	httpmock.Deactivate()
	realURL := "http://129.187.110.174:3000"
	resp, err := http.Get(realURL)
	if err != nil {
		t.Error("Test api not working.")
	}
	fmt.Print(resp)
}

func TestIsAllowedAt(t *testing.T) {
	hyper := HyperLedger{}
	now, _ := time.Parse(time.RFC3339, "2018-01-01T00:00:00Z")
	start, _ := time.Parse(time.RFC3339, "2017-01-01T00:00:00Z")
	end, _ := time.Parse(time.RFC3339, "2019-01-01T00:00:00Z")
	allowed := hyper.isAllowedAt(now, start, end)
	if !allowed {
		t.Error("isAllowedAt not working.")
	}
}

func TestIsNotAllowedAt(t *testing.T) {
	hyper := HyperLedger{}
	now, _ := time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
	start, _ := time.Parse(time.RFC3339, "2017-01-01T00:00:00Z")
	end, _ := time.Parse(time.RFC3339, "2019-01-01T00:00:00Z")
	allowed := hyper.isAllowedAt(now, start, end)
	if allowed {
		t.Error("isAllowedAt not working.")
	}
}
