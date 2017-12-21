package journal
import "time"

type WingLeave struct {
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e WingLeave) GetEvent() string {
       return e.Event
}

func (e WingLeave) GetTimestamp() time.Time {
       return e.Timestamp
}
