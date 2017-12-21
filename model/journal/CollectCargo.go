package journal
import "time"

type CollectCargo struct {
	Stolen    bool   `json:"Stolen"`
	Type      string `json:"Type"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e CollectCargo) GetEvent() string {
       return e.Event
}

func (e CollectCargo) GetTimestamp() time.Time {
       return e.Timestamp
}
