package journal
import "time"

type EngineerProgress struct {
	Engineer  string `json:"Engineer"`
	Progress  string `json:"Progress"`
	Rank      int64  `json:"Rank"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e EngineerProgress) GetEvent() string {
       return e.Event
}

func (e EngineerProgress) GetTimestamp() time.Time {
       return e.Timestamp
}
