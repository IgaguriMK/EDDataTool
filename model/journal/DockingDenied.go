package journal
import "time"

type DockingDenied struct {
	Reason      string `json:"Reason"`
	StationName string `json:"StationName"`
	Event       string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e DockingDenied) GetEvent() string {
       return e.Event
}

func (e DockingDenied) GetTimestamp() time.Time {
       return e.Timestamp
}
