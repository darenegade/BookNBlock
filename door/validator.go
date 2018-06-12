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
		IsAllowedAt(mieter RenterPubkey, t time.Time) (bool, error)
	}
	Lock interface {
		Open()
	}
)

func (v *Validator) Handle(m OpenDoorMessage) error {
	now := time.Now()
	ok, err := v.ContractInfoer.IsAllowedAt(m.RenterPubkey, now)
	if err != nil {
		log.Println("Failed to confirm contract.")
		return err
	}
	if !ok {
		log.Printf("Mieter %s was denied to open Door.\n", m.RenterPubkey)
		return nil
	}

	log.Println("Renter is allowed, opening door.")
	v.Lock.Open()
	return nil
}
