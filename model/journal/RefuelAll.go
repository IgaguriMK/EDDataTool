package journal
import "time"

type RefuelAll struct {
	Amount    float64 `json:"Amount"`
	Cost      int64   `json:"Cost"`
	Event     string  `json:"event"`
	Timestamp time.Time  `json:"timestamp"`
}
func (e RefuelAll) GetEvent() string {
       return e.Event
}

func (e RefuelAll) GetTimestamp() time.Time {
       return e.Timestamp
}
