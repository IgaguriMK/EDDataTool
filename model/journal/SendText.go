package journal
import "time"

type SendText struct {
	Message   string `json:"Message"`
	To        string `json:"To"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e SendText) GetEvent() string {
       return e.Event
}

func (e SendText) GetTimestamp() time.Time {
       return e.Timestamp
}
