package contract

import "testing"
import "fmt"

func TestCreateConnection(t *testing.T) {
	ethCon := new(Ethereum)
	fmt.Print("Run test.")
	ethCon.setPath("https://rinkeby.infura.io/VhXic4UDRfv5w86p2hq7")
	ethCon.setContractAddress("0xEe86D8d8163844517676C918556CDf42310c1671")
	var connected = ethCon.createConnection()
	if !connected {
		t.Error("Testing the connection failed.")
	}
}
