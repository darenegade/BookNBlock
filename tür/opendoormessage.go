package tür

type (
	OpenDoorMessage struct {
		DoorID    DoorID
		RenterPK  RenterPK
		RenterID  RenterID
		Timestamp int
	}
)
