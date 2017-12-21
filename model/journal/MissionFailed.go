package journal
import "time"

type MissionFailed struct {
	MissionID int64  `json:"MissionID"`
	Name      string `json:"Name"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e MissionFailed) GetEvent() string {
       return e.Event
}

func (e MissionFailed) GetTimestamp() time.Time {
       return e.Timestamp
}
