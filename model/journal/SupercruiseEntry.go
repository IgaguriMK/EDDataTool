package journal
import "time"

type SupercruiseEntry struct {
	StarSystem string `json:"StarSystem"`
	Event      string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e SupercruiseEntry) GetEvent() string {
       return e.Event
}

func (e SupercruiseEntry) GetTimestamp() time.Time {
       return e.Timestamp
}
