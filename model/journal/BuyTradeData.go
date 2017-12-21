package journal
import "time"

type BuyTradeData struct {
	Cost      int64  `json:"Cost"`
	System    string `json:"System"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e BuyTradeData) GetEvent() string {
       return e.Event
}

func (e BuyTradeData) GetTimestamp() time.Time {
       return e.Timestamp
}
