package contract

import (
	"fmt"
	"log"
	"math/big"

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

func (e *Ethereum) IsAllowedAt(bookingID *big.Int, renterID common.Address, reqTime *big.Int) (allowed bool) {
	//Annahme: Es existiert ein Contract mit der Methode 'isAllowed(doorID string, renterID string, time time.Time)'
	if e.contract == nil {
		fmt.Print("Contract not initialized yet.")
		return
	}
	callOpts := bind.CallOpts{Pending: true}
	allowed, err := e.contract.IsAllowedAt(&callOpts, bookingID, renterID, reqTime)
	if err != nil {
		fmt.Print("RPC-Call isAllowedAt did not work.")
		return
	}
	return
}

func (e *Ethereum) setPath(path string) {
	e.gethPath = path
}
func (e *Ethereum) setContractAddress(contractAddress string) {
	e.contractAddress = contractAddress
}
func (e *Ethereum) createConnection() (connected bool) {
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
