package message

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/darenegade/BookNBlock/door"
	"github.com/ethereum/go-ethereum/p2p/discover"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	elog "github.com/ethereum/go-ethereum/log"

	"github.com/ethereum/go-ethereum/node"

	"github.com/ethereum/go-ethereum/cmd/utils"

	whisper "github.com/ethereum/go-ethereum/whisper/whisperv6"
)

var (
	// TopicBookNBlock is the topic for sharing BookNBlock messages.
	// It is needed to limit the amount of messages to check.
	TopicBookNBlock = []byte{0x42, 0x6f, 0x6f, 0x6b}
)

type (
	// Whisper is the node that handles the messaging. It uses an ethereum p2p node
	// for networking and the whisper protocol for actual tranfer of messages.
	Whisper struct {
		shh  *whisper.Whisper
		Node *node.Node
	}

	// WhisperConfig holds configuration data for a Whisper. It is needed to
	// define the networking.
	WhisperConfig struct {
		BootstrapNodes []string
		ListenAddr     string
		NodeID         string

		HTTPPort int
	}
)

// StartNode starts an ethereum p2p node and the whisper protocol. Needed to
// start listening for messages. Or providing an api to transmit messages.
func StartNode(config WhisperConfig) *Whisper {
	elog.Root().SetHandler(elog.LvlFilterHandler(elog.Lvl(elog.LvlInfo), elog.StreamHandler(os.Stderr, elog.TerminalFormat(false))))

	nodeConfig := node.DefaultConfig

	nodeConfig.DataDir = "."

	nodeConfig.HTTPHost = "127.0.0.1"
	nodeConfig.HTTPPort = config.HTTPPort
	nodeConfig.HTTPModules = []string{"shh", "admin"}
	nodeConfig.HTTPCors = []string{"*"}
	nodeConfig.HTTPVirtualHosts = []string{"*"}

	nodeConfig.P2P.PrivateKey = config.privateKey()
	nodeConfig.P2P.ListenAddr = config.ListenAddr
	nodeConfig.P2P.BootstrapNodes = config.bootNodes()

	n, err := node.New(&nodeConfig)
	if err != nil {
		utils.Fatalf("Failed to create the protocol node: %v", err)
	}

	whisperConfig := &whisper.DefaultConfig
	w := whisper.New(whisperConfig)

	if err := n.Register(func(n *node.ServiceContext) (node.Service, error) {
		return w, nil
	}); err != nil {
		log.Fatalf("Failed to register the Whisper service: %v", err)
	}

	// start the stack and watch for SIG events
	utils.StartNode(n)

	return &Whisper{
		Node: n,
		shh:  w,
	}
}

// Subscribe to OpenDoorMessages sent to the door (defined by private key).
// Messages that land in the channel are proofen to be from the real sender.
func (w *Whisper) Subscribe(doorPrivateKey door.DoorPrivateKey) (<-chan door.OpenDoorMessage, error) {
	privateKey, err := crypto.HexToECDSA(string(doorPrivateKey))
	if err != nil {
		log.Fatalf("failed to parse private key : %s, error: %s", doorPrivateKey, err)
	}
	filter := whisper.Filter{
		PoW:      whisper.DefaultMinimumPoW,
		Messages: make(map[common.Hash]*whisper.ReceivedMessage),
		AllowP2P: true,
		Topics:   [][]byte{TopicBookNBlock},
		KeyAsym:  privateKey,
	}

	id, err := w.shh.Subscribe(&filter)
	if err != nil {
		panic(err)
	}

	c := make(chan door.OpenDoorMessage)
	go func() {
		ticker := time.NewTicker(250 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if filter := w.shh.GetFilter(id); filter != nil {
					for _, rpcMessage := range filter.Retrieve() {
						ok := rpcMessage.ValidateAndParse()
						if !ok {
							log.Println("Failed to validate rpcMessage, ignoring it.")
							continue // with next
						}
						var openDoorMessage door.OpenDoorMessage
						err := json.Unmarshal(rpcMessage.Payload, &openDoorMessage)
						if err != nil {
							log.Printf("Failed to parse OpenDoorMessage from json, was: %s, err: %s", string(rpcMessage.Payload), err)
							continue // with next
						}
						// fill remaining fields

						c <- openDoorMessage
					}
				}
			}
		}
	}()

	return c, nil
}

func (w Whisper) Post(openDoorMessage door.OpenDoorMessage, privateKey *ecdsa.PrivateKey) error {
	payload, err := json.Marshal(openDoorMessage)
	if err != nil {
		return err
	}
	data, err := hex.DecodeString(string(openDoorMessage.DoorID))
	if err != nil {
		return err
	}
	destination, err := crypto.DecompressPubkey(data)	
	if err != nil {
		return err
	}
	params := &whisper.MessageParams{
		Dst:      destination,
		Payload:  payload,
		PoW:      whisper.DefaultMinimumPoW,
		Src:      privateKey,
		Topic:    whisper.BytesToTopic(TopicBookNBlock),
		TTL:      whisper.DefaultTTL,
		WorkTime: 1,
	}
	message, err := whisper.NewSentMessage(params)
	if err != nil {
		return err
	}
	envelope, err := message.Wrap(params)
	if err != nil {
		return err
	}
	return w.shh.Send(envelope)
}

func (c WhisperConfig) privateKey() *ecdsa.PrivateKey {
	if c.NodeID != "" {
		key, err := crypto.HexToECDSA(c.NodeID)
		if err != nil {
			log.Fatalf("Could not parse private key: %s", err)
		}
		return key
	}
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Could not generate private key: %s", err)
	}
	return key
}

func (c WhisperConfig) bootNodes() (nodes []*discover.Node) {
	for _, node := range c.BootstrapNodes {
		n, err := discover.ParseNode(node)
		if err != nil {
			log.Fatalf("Could not parse node %s: %s", node, err)
		}
		nodes = append(nodes, n)
	}
	return nodes
}
