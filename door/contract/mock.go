package contract

import (
	"time"

	".."
)

type (
	Mock struct {
	}
)

func (Mock) IsAllowedAt(mieter tür.MieterID, t time.Time) (bool, error) {
	return true, nil
}
