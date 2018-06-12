package door

type (
	RenterID string
	RenterPubkey string
	// DoorID is used to identify the door. Might be the doors public key
	DoorID string
	DoorPrivateKey string

	Door struct {
		PrivateKey DoorPrivateKey
		PublicKey  DoorID
	}
)
