package journal
import "time"

type CommitCrime struct {
	CrimeType string `json:"CrimeType"`
	Faction   string `json:"Faction"`
	Fine      int64  `json:"Fine"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e CommitCrime) GetEvent() string {
       return e.Event
}

func (e CommitCrime) GetTimestamp() time.Time {
       return e.Timestamp
}
