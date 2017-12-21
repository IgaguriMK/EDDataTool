package journal
import "time"

type Undocked struct {
	StationName string `json:"StationName"`
	StationType string `json:"StationType"`
	Event       string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e Undocked) GetEvent() string {
       return e.Event
}

func (e Undocked) GetTimestamp() time.Time {
       return e.Timestamp
}
