package journal
import "time"

type ApproachSettlement struct {
	Name      string `json:"Name"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e ApproachSettlement) GetEvent() string {
       return e.Event
}

func (e ApproachSettlement) GetTimestamp() time.Time {
       return e.Timestamp
}
