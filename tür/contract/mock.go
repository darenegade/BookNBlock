package contract

import time "time"

type (
	Mock struct {
	}
)

func (Mock) IsAllowedAt(mieter string, t time.Time) (bool, error) {
	return true, nil
}
