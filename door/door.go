package door

type (
	RenterPrivateKey string
	RenterPublicKey string
	// DoorPublicKey is used to identify the door. Might be the doors public key
	DoorPublicKey string
	DoorPrivateKey string

	BookingID int

	Door struct {
		PrivateKey DoorPrivateKey
		PublicKey  DoorPublicKey
	}
)