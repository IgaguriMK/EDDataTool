package journal
import "time"

type MissionAbandoned struct {
	MissionID int64  `json:"MissionID"`
	Name      string `json:"Name"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e MissionAbandoned) GetEvent() string {
       return e.Event
}

func (e MissionAbandoned) GetTimestamp() time.Time {
       return e.Timestamp
}
