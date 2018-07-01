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
	ethCon.SetContractAddress("0x86f7c1a9fc3a143eB85F2d24Fe3bbE8f8A69dB9c")
	var connected = ethCon.CreateConnection()
	if !connected {
		t.Error("Testing the connection failed.")
	}
}

func TestIsAllowed(t *testing.T) {
	ethCon.SetPath("https://rinkeby.infura.io/VhXic4UDRfv5w86p2hq7")
	ethCon.SetContractAddress("0x86f7c1a9fc3a143eB85F2d24Fe3bbE8f8A69dB9c")
	ethCon.CreateConnection()
	//currentTime := big.NewInt(time.Date(2018, time.May, 30, 0, 0, 0, 0, time.UTC).Unix())
	allowed, err := ethCon.IsAllowedAt(door.BookingID(0), door.RenterPublicKey("0xADF900e582b34EC29DF534e32db6250cf9529FB9"), 1530057600001)
	if err != nil {
		log.Fatalf("%v", err)
	}
	if !allowed {
		t.Error("IsAllowed method failed.")
	}

}

func TestIsNotllowed(t *testing.T) {
	ethCon.SetPath("https://rinkeby.infura.io/VhXic4UDRfv5w86p2hq7")
	ethCon.SetContractAddress("0x86f7c1a9fc3a143eB85F2d24Fe3bbE8f8A69dB9c")
	ethCon.CreateConnection()
	//currentTime := big.NewInt(time.Date(2018, time.May, 30, 0, 0, 0, 0, time.UTC).Unix())
	allowed, err := ethCon.IsAllowedAt(door.BookingID(0), door.RenterPublicKey("0xADF900e582b34EC29DF534e32db6250cf9529FB9"), 1530057500001)
	if err != nil {
		log.Fatalf("%v", err)
	}
	if allowed {
		t.Error("IsAllowed method failed.")
	}

}
