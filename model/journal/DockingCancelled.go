package journal
import "time"

type DockingCancelled struct {
	StationName string `json:"StationName"`
	Event       string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e DockingCancelled) GetEvent() string {
       return e.Event
}

func (e DockingCancelled) GetTimestamp() time.Time {
       return e.Timestamp
}
