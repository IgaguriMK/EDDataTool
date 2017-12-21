package journal
import "time"

type MarketSell struct {
	AvgPricePaid int64  `json:"AvgPricePaid"`
	BlackMarket  bool   `json:"BlackMarket"`
	Count        int64  `json:"Count"`
	IllegalGoods bool   `json:"IllegalGoods"`
	SellPrice    int64  `json:"SellPrice"`
	StolenGoods  bool   `json:"StolenGoods"`
	TotalSale    int64  `json:"TotalSale"`
	Type         string `json:"Type"`
	Event        string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e MarketSell) GetEvent() string {
       return e.Event
}

func (e MarketSell) GetTimestamp() time.Time {
       return e.Timestamp
}
