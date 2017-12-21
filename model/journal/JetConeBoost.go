package journal

import "time"

type JetConeBoost struct {
	BoostValue float64   `json:"BoostValue"`
	Event      string    `json:"event"`
	Timestamp  time.Time `json:"timestamp"`
}

func (e JetConeBoost) GetEvent() string {
	return e.Event
}

func (e JetConeBoost) GetTimestamp() time.Time {
	return e.Timestamp
}
