package door

import (
	"log"
	"time"
)

type (
	Validator struct {
		ContractInfoer
		Lock
	}

	ContractInfoer interface {
		IsAllowedAt(booking BookingID, renter RenterPublicKey, timestamp int64)  (bool, error)
	}
	Lock interface {
		Open()
	}
)

func (v *Validator) Handle(m OpenDoorMessage) error {
	now := time.Now()
	ok, err := v.ContractInfoer.IsAllowedAt(m.Booking, m.RenterPubkey, now.UnixNano())
	if err != nil {
		log.Println("Failed to confirm contract.")
		return err
	}
	if !ok {
		log.Printf("Mieter %s was denied to open Door.\n", m.RenterPubkey)
		log.Printf("Time: %v, TimeUnix: %d, Message: %v\n", now, now.UnixNano(), m)
		return nil
	}

	log.Println("Renter is allowed, opening door.")
	v.Lock.Open()
	return nil
}
