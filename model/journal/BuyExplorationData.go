package journal
import "time"

type BuyExplorationData struct {
	Cost      int64  `json:"Cost"`
	System    string `json:"System"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e BuyExplorationData) GetEvent() string {
       return e.Event
}

func (e BuyExplorationData) GetTimestamp() time.Time {
       return e.Timestamp
}
