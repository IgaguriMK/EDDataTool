package journal
import "time"

type FuelScoop struct {
	Scooped   float64 `json:"Scooped"`
	Total     float64 `json:"Total"`
	Event     string  `json:"event"`
	Timestamp time.Time  `json:"timestamp"`
}
func (e FuelScoop) GetEvent() string {
       return e.Event
}

func (e FuelScoop) GetTimestamp() time.Time {
       return e.Timestamp
}
