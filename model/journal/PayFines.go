package journal
import "time"

type PayFines struct {
	Amount    int64  `json:"Amount"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e PayFines) GetEvent() string {
       return e.Event
}

func (e PayFines) GetTimestamp() time.Time {
       return e.Timestamp
}
