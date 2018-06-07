package door

import "time"

type (
	OpenDoorMessage struct {
		DoorPublicKey   DoorPublicKey   `json:"doorPublicKey"`
		Timestamp       time.Time       `json:"timestamp"`
		RenterPublicKey RenterPublicKey `json:"renterPublicKey"`
	}
)
