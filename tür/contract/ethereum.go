package contract

import (
	"time"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"

	".."
)

type (
	Ethereum struct {
		//Pfad zur GETH Node
		gethPath string
		//Adresse zum Contract
		contractAddress string
		//Instanz des schon im Network deployten Contract
		contract LockContract
	}
)

func (e *Ethereum) IsAllowedAt(mieter tür.MieterID, t time.Time) (allowed bool, error) {
	//Annahme: Es existiert ein Contract mit der Methode 'isAllowed(mieterId string, time time.Time)'
	allowed = e.contract.isAllowedAt(mieter,t)
}
func (e *Ethereum) setPath(path string){
	e.path=path
}
func (e *Ethereum) setContractAddress(contractAddress string){
	e.contractAddress=contractAddress
}
func (e *Ethererum) createConnection(){
	//  IPC basierte RPC-Verbindung wird zur entfernten Node hergestellt
	conn, err := ethclient.Dial(e.gethPath)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	//Annahme: Contract mit Namen LockContract wurde schon zum Network deployed
	e.contract, err := LockContract(common.HexToAddress(e.contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
}
func main(){
	ethCon := Ethereum
	ethCon.setPath("https://rinkeby.infura.io/VhXic4UDRfv5w86p2hq7")
	ethCon.setContractAddress("0xEe86D8d8163844517676C918556CDf42310c1671")
	ehtCon.createConnection()
	
}
}
