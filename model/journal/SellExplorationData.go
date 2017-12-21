package journal
import "time"

type SellExplorationData struct {
	BaseValue  int64    `json:"BaseValue"`
	Bonus      int64    `json:"Bonus"`
	Discovered []string `json:"Discovered"`
	Systems    []string `json:"Systems"`
	Event      string   `json:"event"`
	Timestamp time.Time   `json:"timestamp"`
}
func (e SellExplorationData) GetEvent() string {
       return e.Event
}

func (e SellExplorationData) GetTimestamp() time.Time {
       return e.Timestamp
}
