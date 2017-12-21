package journal
import "time"

type Resurrect struct {
	Bankrupt  bool   `json:"Bankrupt"`
	Cost      int64  `json:"Cost"`
	Option    string `json:"Option"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e Resurrect) GetEvent() string {
       return e.Event
}

func (e Resurrect) GetTimestamp() time.Time {
       return e.Timestamp
}
