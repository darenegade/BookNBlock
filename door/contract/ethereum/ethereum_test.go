package contract

import (
	"fmt"
	"os"
	"testing"
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

	ethCon.setPath("https://rinkeby.infura.io/VhXic4UDRfv5w86p2hq7")
	ethCon.setContractAddress("0xEe86D8d8163844517676C918556CDf42310c1671")
	var connected = ethCon.createConnection()
	if !connected {
		t.Error("Testing the connection failed.")
	}
}

// func TestIsAllowed(t *testing.T) {
// 	ethCon.setPath("https://rinkeby.infura.io/VhXic4UDRfv5w86p2hq7")
// 	ethCon.setContractAddress("0xEe86D8d8163844517676C918556CDf42310c1671")
// 	ethCon.createConnection()
// 	var allowed = ethCon.IsAllowedAt("Franz", time.Parse(time.RFC822, "01 Jan 15 10:00 UTC"))
// 	if !allowed {
// 		t.Error("IsAllowed method failed.")
// 	}

// }
