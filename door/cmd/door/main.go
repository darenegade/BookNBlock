package main

import (
	"flag"
	"log"

	"github.com/darenegade/BookNBlock/door"
	"github.com/darenegade/BookNBlock/door/contract/ethereum"
	"github.com/darenegade/BookNBlock/door/lock"
	"github.com/darenegade/BookNBlock/door/logserver"
	"github.com/darenegade/BookNBlock/door/message"
)

type (
	Config struct {
		BootstrapNode string
		PrivateKey    string
		LogServer     bool
		ListenAddr    string
		HTTPHost      string
		HTTPPort      int
	}
)

func main() {
	defer lock.Lock.Finish()
	var config Config
	flag.StringVar(&config.PrivateKey, "key", "6ecd6756d5e9d9df44be83b82d99b17983ae5ce9d0f2de9dcd68c80197aafc4a", "private key to proof that this the the door it pretends to be")
	flag.StringVar(&config.BootstrapNode, "bootnode", "enode://7d13360f5b1ddcf6947f244639113597a863abba0589d2fa5fffb2816ead0acea6211d5778a8be648e45e81ed881f4c1f5c9bbbf0e79065dfb54bcd97de3beab@127.0.0.1:8066", "bootnode to start connection to the ethereum network")
	flag.BoolVar(&config.LogServer, "logserver", false, "enable the logging server")
	flag.StringVar(&config.ListenAddr, "listenaddr", "127.0.0.1:8067", "listen address for other nodes")
	flag.IntVar(&config.HTTPPort, "httpport", 9945, "http port for api")
	flag.StringVar(&config.HTTPHost, "httphost", "127.0.0.1", "http host for api")
	flag.Parse()

	whisperConfig := message.WhisperConfig{
		BootstrapNodes: []string{config.BootstrapNode},
		NodeID:         config.PrivateKey,
		ListenAddr:     config.ListenAddr,
		HTTPHost:       config.HTTPHost,
		HTTPPort:       config.HTTPPort,
	}

	if config.LogServer {
		go logging()
	}

	var ethCon contract.Ethereum
	whisper := message.StartNode(whisperConfig)
	ethCon.SetPath("https://rinkeby.infura.io/VhXic4UDRfv5w86p2hq7")
	ethCon.SetContractAddress("0x5A6d0aD02efb81C6c4dd3790f33f410dB5c8cD0b")
	var connected = ethCon.CreateConnection()
	if !connected {
		log.Fatal("Connection to ethereum failed.")
	}

	c, err := whisper.Subscribe(door.DoorPrivateKey(config.PrivateKey))
	if err != nil {
		log.Fatal(err)
	}

	validator := door.Validator{
		ContractInfoer: &ethCon,
		Lock:           lock.Lock,
	}

	for m := range c {
		validator.Handle(m)
	}
}

func logging() {
	logserver.SetLogging("log.txt")
	logserver.StartWebserver()
}
