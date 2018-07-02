package main

import (
	"github.com/darenegade/BookNBlock/door/message"
	"github.com/darenegade/BookNBlock/door"
	"log"
	"github.com/darenegade/BookNBlock/door/contract/ethereum"
	"flag"
	"github.com/darenegade/BookNBlock/door/lock"
	"github.com/darenegade/BookNBlock/door/logserver"
)

type (
	Config struct {
		BootstrapNode string
		PrivateKey string
		LogServer bool
	}
)

func main() {
	var config Config
	flag.StringVar(&config.PrivateKey, "key", "6ecd6756d5e9d9df44be83b82d99b17983ae5ce9d0f2de9dcd68c80197aafc4a", "private key to proof that this the the door it pretends to be")
	flag.StringVar(&config.BootstrapNode, "bootnode", "enode://7d13360f5b1ddcf6947f244639113597a863abba0589d2fa5fffb2816ead0acea6211d5778a8be648e45e81ed881f4c1f5c9bbbf0e79065dfb54bcd97de3beab@127.0.0.1:30349", "bootnode to start connection to the ethereum network")
	flag.BoolVar(&config.LogServer, "logserver", false,"enable the logging server")
	flag.Parse()

	whisperConfig := message.WhisperConfig{
		BootstrapNodes: []string{config.BootstrapNode},
		NodeID:         config.PrivateKey,
	}

	if config.LogServer {
		go logging()
	}


	var ethCon contract.Ethereum
	whisper := message.StartNode(whisperConfig)
	ethCon.SetPath("https://rinkeby.infura.io/VhXic4UDRfv5w86p2hq7")
	ethCon.SetContractAddress("0x86f7c1a9fc3a143eB85F2d24Fe3bbE8f8A69dB9c")
	var connected = ethCon.CreateConnection()
	if !connected {
		log.Fatal("Connection to ethereum failed.")
	}

	c,err  := whisper.Subscribe(door.DoorPrivateKey(config.PrivateKey))
	if err != nil {
		log.Fatal(err)
	}

	validator := door.Validator{
		ContractInfoer: &ethCon,
		Lock: lock.Lock,
	}

	for m := range c {
		validator.Handle(m)
	}
}

func logging() {
	logserver.SetLogging("log.txt")
	logserver.StartWebserver()
}