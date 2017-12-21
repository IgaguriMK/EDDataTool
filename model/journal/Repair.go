package journal
import "time"

type Repair struct {
	Cost      int64  `json:"Cost"`
	Item      string `json:"Item"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e Repair) GetEvent() string {
       return e.Event
}

func (e Repair) GetTimestamp() time.Time {
       return e.Timestamp
}
