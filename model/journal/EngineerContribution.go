package journal
import "time"

type EngineerContribution struct {
	Commodity     string `json:"Commodity"`
	Engineer      string `json:"Engineer"`
	Quantity      int64  `json:"Quantity"`
	TotalQuantity int64  `json:"TotalQuantity"`
	Type          string `json:"Type"`
	Event         string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e EngineerContribution) GetEvent() string {
       return e.Event
}

func (e EngineerContribution) GetTimestamp() time.Time {
       return e.Timestamp
}
