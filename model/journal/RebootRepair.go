package journal
import "time"

type RebootRepair struct {
	Modules   []string `json:"Modules"`
	Event     string   `json:"event"`
	Timestamp time.Time   `json:"timestamp"`
}
func (e RebootRepair) GetEvent() string {
       return e.Event
}

func (e RebootRepair) GetTimestamp() time.Time {
       return e.Timestamp
}
