package journal
import "time"

type BuyAmmo struct {
	Cost      int64  `json:"Cost"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e BuyAmmo) GetEvent() string {
       return e.Event
}

func (e BuyAmmo) GetTimestamp() time.Time {
       return e.Timestamp
}
