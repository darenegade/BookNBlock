package main

import (
	"time"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"fmt"

)

type Ethereum struct {
		//Pfad zur GETH Node
		gethPath string
		//Adresse zum Contract
		contractAddress string
		//Instanz des schon deployten Contracts
		contract LockContract
		//Id der dem Contract zugeordneten Tuer
		doorID string		
		
}

func (e *Ethereum) IsAllowedAt(renter door.RenterID, t time.Time) (allowed bool) {
	//Annahme: Es existiert ein Contract mit der Methode 'isAllowed(doorID string, renterID string, time time.Time)'
	allowed = e.contract.isAllowedAt(e.doorID, renter,t)
}
func (e *Ethereum) setDoorID(doorID string){
	e.doorID=doorID
}
func (e *Ethereum) setPath(path string){
	e.gethPath=path
}
func (e *Ethereum) setContractAddress(contractAddress string){
	e.contractAddress=contractAddress
}
func (e *Ethereum) createConnection(){
	//  IPC basierte RPC-Verbindung wird zur entfernten Node hergestellt
	conn, err := ethclient.Dial(e.gethPath)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	//Annahme: Contract mit Namen LockContract wurde schon zum Network deployed
	e.contract, err := NewLockContract(common.HexToAddress(e.contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
}
	//fmt.Print(contract)
}

func main(){
	ethCon := Ethereum{gethPath:"https://rinkeby.infura.io/VhXic4UDRfv5w86p2hq7", contractAddress:"0xEe86D8d8163844517676C918556CDf42310c1671"}
	// ethCon.setPath("https://rinkeby.infura.io/VhXic4UDRfv5w86p2hq7")
	// ethCon.setContractAddress("0xEe86D8d8163844517676C918556CDf42310c1671")
	ethCon.createConnection()
	//ethCon.isAllowedAt()
	

}

