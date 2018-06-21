package door

import (
	"github.com/darenegade/BookNBlock/door/message"
	"github.com/darenegade/BookNBlock/door"
	"log"
	"github.com/darenegade/BookNBlock/door/contract/ethereum"
)

type (
	Config struct {
		BootstrapNode string
		PrivateKey string
	}
)

func main() {
	var config Config
	whisperConfig := message.WhisperConfig{
		BootstrapNodes: []string{config.BootstrapNode},
		NodeID:         config.PrivateKey,
	}


	var contract *contract.Ethereum
	whisper := message.StartNode(whisperConfig)
	c,err  := whisper.Subscribe(door.DoorPrivateKey(config.PrivateKey))
	if err != nil {
		log.Fatal(err)
	}

	validator := door.Validator{
		ContractInfoer: contract,
	}

	for m := range c {
		validator.Handle(m)
	}
}