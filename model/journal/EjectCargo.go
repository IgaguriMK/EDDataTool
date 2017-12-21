package journal
import "time"

type EjectCargo struct {
	Abandoned bool   `json:"Abandoned"`
	Count     int64  `json:"Count"`
	Type      string `json:"Type"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e EjectCargo) GetEvent() string {
       return e.Event
}

func (e EjectCargo) GetTimestamp() time.Time {
       return e.Timestamp
}
