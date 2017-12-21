package journal
import "time"

type HeatWarning struct {
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e HeatWarning) GetEvent() string {
       return e.Event
}

func (e HeatWarning) GetTimestamp() time.Time {
       return e.Timestamp
}
