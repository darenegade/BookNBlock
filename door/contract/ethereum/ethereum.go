package contract

import (
	"fmt"
	"log"
	"math/big"

	"github.com/darenegade/BookNBlock/door"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Ethereum struct {
	//Pfad zur GETH Node
	gethPath string
	//Adresse zum Contract
	contractAddress string
	//Instanz des schon deployten Contracts
	contract *LockContract
	//Die dem contract zugordnete Tür
	offer Offer
}

//Zugeordnetes Angebot hat eine eindeutige ID und ist einem Vermieter und einer Tür zugeordnet
type Offer struct {
	OfferID    string
	LandlordID string
	RenterID   string
	DoorID     string
}

func (e *Ethereum) IsAllowedAt(booking door.BookingID, renter door.RenterPublicKey, timestamp int) (allowed bool, err error) {
	//Annahme: Es existiert ein Contract mit der Methode 'isAllowed(doorID string, renterID string, time time.Time)'
	if e.contract == nil {
		fmt.Print("Contract not initialized yet.")
		return
	}
	callOpts := bind.CallOpts{Pending: true}

	allowed, err = e.contract.IsAllowedAt(&callOpts, big.NewInt(int64(booking)), common.HexToAddress(string(renter)), big.NewInt(int64(timestamp)))
	// fmt.Println(allowed)
	// allowed = true
	// test2, err := e.contract.GetOfferIDs(&callOpts)
	// fmt.Println()
	// fmt.Println(test2[0])
	if err != nil {
		log.Fatalf("RPC-Call isAllowedAt did not work: %v", err)
	}
	return allowed, nil
}

func (e *Ethereum) SetPath(path string) {
	e.gethPath = path
}
func (e *Ethereum) SetContractAddress(contractAddress string) {
	e.contractAddress = contractAddress
}
func (e *Ethereum) CreateConnection() (connected bool) {
	//  IPC basierte RPC-Verbindung wird zur entfernten Node hergestellt
	conn, err := ethclient.Dial(e.gethPath)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		return false
	}
	//Annahme: Contract mit Namen LockContract wurde schon zum Network deployed
	contract, err := NewLockContract(common.HexToAddress(e.contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
		return false
	}
	e.contract = contract
	return true
}
