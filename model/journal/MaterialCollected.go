package journal
import "time"

type MaterialCollected struct {
	Category  string `json:"Category"`
	Count     int64  `json:"Count"`
	Name      string `json:"Name"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e MaterialCollected) GetEvent() string {
       return e.Event
}

func (e MaterialCollected) GetTimestamp() time.Time {
       return e.Timestamp
}
