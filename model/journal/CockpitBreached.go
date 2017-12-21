package journal
import "time"

type CockpitBreached struct {
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e CockpitBreached) GetEvent() string {
       return e.Event
}

func (e CockpitBreached) GetTimestamp() time.Time {
       return e.Timestamp
}
