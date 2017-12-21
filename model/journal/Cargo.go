package journal
import "time"

type Cargo struct {
	Inventory []struct {
		Count int64  `json:"Count"`
		Name  string `json:"Name"`
	} `json:"Inventory"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e Cargo) GetEvent() string {
       return e.Event
}

func (e Cargo) GetTimestamp() time.Time {
       return e.Timestamp
}
