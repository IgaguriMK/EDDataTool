package journal
import "time"

type DockingRequested struct {
	StationName string `json:"StationName"`
	Event       string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e DockingRequested) GetEvent() string {
       return e.Event
}

func (e DockingRequested) GetTimestamp() time.Time {
       return e.Timestamp
}
