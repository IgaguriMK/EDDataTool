package journal
import "time"

type BuyDrones struct {
	BuyPrice  int64  `json:"BuyPrice"`
	Count     int64  `json:"Count"`
	TotalCost int64  `json:"TotalCost"`
	Type      string `json:"Type"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e BuyDrones) GetEvent() string {
       return e.Event
}

func (e BuyDrones) GetTimestamp() time.Time {
       return e.Timestamp
}
