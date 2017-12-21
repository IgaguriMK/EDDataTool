package journal
import "time"

type NavBeaconScan struct {
	NumBodies int64  `json:"NumBodies"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e NavBeaconScan) GetEvent() string {
       return e.Event
}

func (e NavBeaconScan) GetTimestamp() time.Time {
       return e.Timestamp
}
