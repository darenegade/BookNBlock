package door

type (
	RenterID string
	RenterPubkey string
	DoorPublicKey string
	DoorPrivateKey string

	Door struct {
		PrivateKey DoorPrivateKey
		PublicKey  DoorPublicKey
	}
)
