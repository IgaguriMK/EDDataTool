package journal
import "time"

type ShieldState struct {
	ShieldsUp bool   `json:"ShieldsUp"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e ShieldState) GetEvent() string {
       return e.Event
}

func (e ShieldState) GetTimestamp() time.Time {
       return e.Timestamp
}
