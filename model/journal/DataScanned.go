package journal
import "time"

type DataScanned struct {
	Type      string `json:"Type"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e DataScanned) GetEvent() string {
       return e.Event
}

func (e DataScanned) GetTimestamp() time.Time {
       return e.Timestamp
}
