package journal
import "time"

type SelfDestruct struct {
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e SelfDestruct) GetEvent() string {
       return e.Event
}

func (e SelfDestruct) GetTimestamp() time.Time {
       return e.Timestamp
}
