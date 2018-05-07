package contract

import (
	"time"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	".."
)

type (
	Ethereum struct {
		//Path to GETH Node
		gethPath string
		//Address to the contract
		contractAddress string
	}
)

func (e *Ethereum) IsAllowedAt(mieter tür.MieterID, t time.Time) (bool, error) {
	panic("not yet implemented")
}
func (e. *Ethereum) setPath(path string){
	e.path=path
}
func (e *Ethererum) createConnection(){
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(e.path)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	contract, err := NewTronToken(common.HexToAddress("0xf230b790E05390FC8295F4d3F60332c93BEd42e2"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
}
}
