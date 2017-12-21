package journal
import "time"

type WingJoin struct {
	Others    []string `json:"Others"`
	Event     string   `json:"event"`
	Timestamp time.Time   `json:"timestamp"`
}
func (e WingJoin) GetEvent() string {
       return e.Event
}

func (e WingJoin) GetTimestamp() time.Time {
       return e.Timestamp
}
