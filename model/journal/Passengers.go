package journal
import "time"

type Passengers struct {
	Manifest []struct {
		Count     int64  `json:"Count"`
		MissionID int64  `json:"MissionID"`
		Type      string `json:"Type"`
		Vip       bool   `json:"VIP"`
		Wanted    bool   `json:"Wanted"`
	} `json:"Manifest"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e Passengers) GetEvent() string {
       return e.Event
}

func (e Passengers) GetTimestamp() time.Time {
       return e.Timestamp
}
