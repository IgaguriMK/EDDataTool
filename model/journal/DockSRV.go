package journal
import "time"

type DockSRV struct {
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e DockSRV) GetEvent() string {
       return e.Event
}

func (e DockSRV) GetTimestamp() time.Time {
       return e.Timestamp
}
