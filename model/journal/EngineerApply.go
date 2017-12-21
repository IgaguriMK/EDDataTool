package journal
import "time"

type EngineerApply struct {
	Blueprint string `json:"Blueprint"`
	Engineer  string `json:"Engineer"`
	Level     int64  `json:"Level"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e EngineerApply) GetEvent() string {
       return e.Event
}

func (e EngineerApply) GetTimestamp() time.Time {
       return e.Timestamp
}
