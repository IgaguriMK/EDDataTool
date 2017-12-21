package journal
import "time"

type Synthesis struct {
	Materials []struct {
		Count int64  `json:"Count"`
		Name  string `json:"Name"`
	} `json:"Materials"`
	Name      string `json:"Name"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e Synthesis) GetEvent() string {
       return e.Event
}

func (e Synthesis) GetTimestamp() time.Time {
       return e.Timestamp
}
