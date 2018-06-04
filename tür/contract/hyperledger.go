package contract

import time "time"

type (
	HyperLedger struct {
		// ...?
	}
)

func (e *HyperLedger) IsAllowedAt(mieter string, t time.Time) (bool, error) {
	panic("not yet implemented")
}
