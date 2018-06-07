package door

type (
	DoorPublicKey   string
	DoorPrivateKey  string
	RenterPublicKey string

	Door struct {
		PublicKey  DoorPublicKey
		PrivateKey DoorPrivateKey
	}
)
