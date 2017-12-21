package journal
import "time"

type MarketBuy struct {
	BuyPrice  int64  `json:"BuyPrice"`
	Count     int64  `json:"Count"`
	TotalCost int64  `json:"TotalCost"`
	Type      string `json:"Type"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e MarketBuy) GetEvent() string {
       return e.Event
}

func (e MarketBuy) GetTimestamp() time.Time {
       return e.Timestamp
}
