package contract

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/darenegade/BookNBlock/door"
)

var ethCon Ethereum

func TestMain(m *testing.M) {
	setUp()
	ethCon := new(Ethereum)
	fmt.Print(ethCon)
	fmt.Print("Run test.")
	retCode := m.Run()
	os.Exit(retCode)
}

func setUp() {

}

func TestCreateConnection(t *testing.T) {

	ethCon.SetPath("https://rinkeby.infura.io/VhXic4UDRfv5w86p2hq7")
	ethCon.SetContractAddress("0xEe86D8d8163844517676C918556CDf42310c1671")
	var connected = ethCon.CreateConnection()
	if !connected {
		t.Error("Testing the connection failed.")
	}
}

func TestIsAllowed(t *testing.T) {
	ethCon.SetPath("https://rinkeby.infura.io/VhXic4UDRfv5w86p2hq7")
	ethCon.SetContractAddress("0xEe86D8d8163844517676C918556CDf42310c1671")
	ethCon.CreateConnection()
	allowed, err := ethCon.IsAllowedAt(door.BookingID(123), door.RenterPublicKey("123"), 4)
	if err != nil {
		log.Fatalf("%v", err)
	}
	if !allowed {
		t.Error("IsAllowed method failed.")
	}

}
