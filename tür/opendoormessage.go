package tür

type (
	OpenDoorMessage struct {
		DoorID       DoorID
		RenterID     RenterID
		RenterPubkey RenterPubkey
		Timestamp    int64
	}
)
