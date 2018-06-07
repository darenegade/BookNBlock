package door

import "time"

type (
	OpenDoorMessage struct {
		DoorPublicKey   DoorPublicKey   `json:"doorPublicKey"`
		BookingID       string          `json:"bookingID"`
		Timestamp       time.Time       `json:"timestamp"`
		RenterPublicKey RenterPublicKey `json:"renterPublicKey"`
	}
)
