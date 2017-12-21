package journal
import "time"

type WingAdd struct {
	Name      string `json:"Name"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e WingAdd) GetEvent() string {
       return e.Event
}

func (e WingAdd) GetTimestamp() time.Time {
       return e.Timestamp
}
