package main

import (
	"github.com/darenegade/BookNBlock/door"
	"github.com/darenegade/BookNBlock/door/message"
	"bufio"
	"os"
	"time"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"fmt"
	"strconv"
	"flag"
)

const RENTER_PRIVATE_KEY =
	`c81803c093599fb7a4bba473ed74a2209fa53fcb59a7a7bcc6e8226157e87cb5`;
const DOOR_PRIVATE_KEY =
	`6ecd6756d5e9d9df44be83b82d99b17983ae5ce9d0f2de9dcd68c80197aafc4a`;
const DOOR_PUBLIC_KEY =
	`04f0f871df7b11b3a186210ef251d10837ccfb757de9d8669225bcf73632853def72ae7680f8acdfa1ac94345017d2b4c185275a1ea2f7bbe03e939146ba355889`;

// main used to start an bootnode for the whisper protocol. It is used to provide an initial
// central network point for door or renter to join.
func main() {
	c := message.WhisperConfig{}
	flag.StringVar(&c.ListenAddr, "listenaddr", "127.0.0.1:30349", "listen address for other nodes")
	flag.StringVar(&c.NodeID, "nodeid", "b3651aff593ef395ee7c16f3ca681830f7d8d0b2729cf472b14f2c4ebe833aa0", "node id is private key for the door")
	flag.IntVar(&c.HTTPPort, "httpport", 9945,"http port for api")
	flag.StringVar(&c.HTTPHost, "httphost", "127.0.0.1","http host for api")
	flag.Parse()

	w := message.StartNode(c)
	mess,_ := w.Subscribe(door.DoorPrivateKey(DOOR_PRIVATE_KEY))
	go func() {
		for m := range mess {
			fmt.Printf("received %#v\n", m)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("input bookinID and press enter")
	for scanner.Scan() {
		booking, _ := strconv.Atoi(scanner.Text())
		opendoormessage := door.OpenDoorMessage{
			DoorID: DOOR_PUBLIC_KEY,
			RenterPubkey: door.RenterPublicKey("0xADF900e582b34EC29DF534e32db6250cf9529FB9"),
			Timestamp: int(time.Now().Unix()),
			Booking: door.BookingID(booking),
		}

		privateKey, err := crypto.HexToECDSA(RENTER_PRIVATE_KEY)
		if err != nil {
			log.Println(err)
		}
		w.Post(opendoormessage,privateKey )
		fmt.Printf("posted with booking %d\n", booking)
	}
	w.Node.Stop()
}
