package journal
import "time"

type DockingGranted struct {
	LandingPad  int64  `json:"LandingPad"`
	StationName string `json:"StationName"`
	Event       string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e DockingGranted) GetEvent() string {
       return e.Event
}

func (e DockingGranted) GetTimestamp() time.Time {
       return e.Timestamp
}
