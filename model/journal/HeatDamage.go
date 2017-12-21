package journal
import "time"

type HeatDamage struct {
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e HeatDamage) GetEvent() string {
       return e.Event
}

func (e HeatDamage) GetTimestamp() time.Time {
       return e.Timestamp
}
